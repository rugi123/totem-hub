package ws

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	lru "github.com/hashicorp/golang-lru"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/internal/dto"
)

type MemberGetter interface {
	GetByID(ctx context.Context, memberID uuid.UUID) (*entity.Member, error)
}

type Client struct {
	entity.Member
	Conn *websocket.Conn
	Send chan dto.BroadcastMessage

	mu     sync.Mutex
	closed bool
}

type WebsocketHub struct {
	mu sync.RWMutex

	memberGetter MemberGetter

	connections   map[uuid.UUID]*Client
	membersByChat map[uuid.UUID]map[uuid.UUID]bool
	membersByUser map[uuid.UUID]map[uuid.UUID]bool
	membersCache  *lru.Cache

	shutdownChan chan bool
	waitGroup    sync.WaitGroup
}

func NewWebsocketHub(memberGetter MemberGetter, cacheSize int) (*WebsocketHub, error) {
	cache, err := lru.New(cacheSize)
	if err != nil {
		return nil, err
	}

	return &WebsocketHub{
		memberGetter:  memberGetter,
		connections:   make(map[uuid.UUID]*Client),
		membersByChat: make(map[uuid.UUID]map[uuid.UUID]bool),
		membersByUser: make(map[uuid.UUID]map[uuid.UUID]bool),
		membersCache:  cache,
		shutdownChan:  make(chan bool),
	}, nil
}

func (h *WebsocketHub) RegisterClient(conn *websocket.Conn, memberID uuid.UUID) error {
	member, err := h.getMember(memberID)
	if err != nil {
		return err
	}

	client := &Client{
		Member: *member,
		Conn:   conn,
		Send:   make(chan dto.BroadcastMessage, 256),
	}
	h.mu.Lock()
	defer h.mu.Unlock()

	if existingClient, exists := h.connections[memberID]; exists {
		closeClient(existingClient)
	}

	h.connections[member.ID] = client

	// Обновляем индексы
	if h.membersByChat[member.ChatID] == nil {
		h.membersByChat[member.ChatID] = make(map[uuid.UUID]bool)
	}
	h.membersByChat[member.ChatID][member.ID] = true

	if h.membersByUser[member.UserID] == nil {
		h.membersByUser[member.UserID] = make(map[uuid.UUID]bool)
	}
	h.membersByUser[member.UserID][member.ID] = true

	go h.writePump(client)
	return nil
}

func (h *WebsocketHub) BroadcastToChat(memberID uuid.UUID, message dto.BroadcastMessage) error {
	member, err := h.getMember(memberID)
	if err != nil {
		return err
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	memberIDs, exists := h.membersByChat[member.ChatID]
	if !exists {
		return nil
	}
	for memberID := range memberIDs {
		client, exists := h.connections[memberID]
		if !exists || client == nil {
			continue
		}

		safeSend(client, message)
	}
	return nil
}

func (h *WebsocketHub) getMember(memberID uuid.UUID) (*entity.Member, error) {
	if cached, ok := h.membersCache.Get(memberID); ok {
		if member, ok := cached.(*entity.Member); ok {
			return member, nil
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	member, err := h.memberGetter.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	h.membersCache.Add(memberID, member)

	return member, nil
}

func closeClient(client *Client) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closed {
		return
	}

	close(client.Send)
	client.Conn.Close()
	client.closed = true
}

func safeSend(client *Client, message dto.BroadcastMessage) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closed {
		return
	}

	select {
	case client.Send <- message:
	default:
		log.Printf("Client %s send channel full, disconnecting", client.ID)
		closeClient(client)
	}
}

func (h *WebsocketHub) writePump(client *Client) {
	defer h.waitGroup.Done()

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				return
			}

			if err := h.writeMessage(client.Conn, message); err != nil {
				log.Printf("Write error: %v", err)
				return
			}
		case <-ticker.C:
			if err := h.ping(client.Conn); err != nil {
				log.Printf("Ping error: %v", err)
				return
			}
		case <-h.shutdownChan:
			return
		}
	}
}

func (h *WebsocketHub) writeMessage(conn *websocket.Conn, message dto.BroadcastMessage) error {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return conn.WriteJSON(message)
}

func (h *WebsocketHub) ping(conn *websocket.Conn) error {
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return conn.WriteMessage(websocket.PingMessage, nil)
}
