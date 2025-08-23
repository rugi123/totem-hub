package websocket

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"golang.org/x/net/websocket"
)

type MemberGetter interface {
	GetByID(ctx context.Context, memberID uuid.UUID) (*entity.Member, error)
}

type Client struct {
	entity.Member
	Conn *websocket.Conn
	Send chan []byte

	mu     sync.Mutex
	closed bool
}

type WebsocketHub struct {
	mu sync.RWMutex

	memberGetter MemberGetter

	connections   map[uuid.UUID]*Client
	membersByChat map[uuid.UUID]map[uuid.UUID]bool
	membersByUser map[uuid.UUID]map[uuid.UUID]bool
}

func (h *WebsocketHub) RegisterClient(conn *websocket.Conn, memberID uuid.UUID) error {
	member, err := h.GetFromCache(memberID)
	if err != nil {
		return err
	}

	client := &Client{
		Member: *member,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}
	h.mu.Lock()
	defer h.mu.Unlock()

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
	return nil
}

func (h *WebsocketHub) BroadcastToChat(memberID uuid.UUID, message string) error {
	member, err := h.GetFromCache(memberID)
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
		if !exists || client == nil || client.IsMuted {
			continue
		}

		client.Send <- []byte(message)
	}
	return nil
}

func (h *WebsocketHub) GetFromCache(memberID uuid.UUID) (*entity.Member, error) {
	member, exists := h.membersCache[memberID]
	if exists {
		return member, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	member, err := h.memberGetter.GetByID(ctx, memberID)

	return member, err
}
