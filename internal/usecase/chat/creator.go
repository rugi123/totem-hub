package chat

import (
	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/internal/dto"
	"github.com/rugi123/totem-hub/pkg/validator"
)

type ChatCreator interface {
	ValidateAttributes(attrs map[string]interface{}) error
	CreateEntity(chatType, title string, userID uuid.UUID, attrs map[string]interface{}) (entity.Chat, error)
}

// Базовый тип
type BaseChatCreator struct{}

func (c *BaseChatCreator) FilterAttributes(attrs map[string]interface{}, allowedKeys []string) map[string]interface{} {
	filtered := make(map[string]interface{})
	for _, key := range allowedKeys {
		if val, ok := attrs[key]; ok {
			filtered[key] = val
		}
	}
	return filtered
}

type ChannelCreator struct {
	BaseChatCreator
}

func (c *ChannelCreator) ValidateAttributes(attrs map[string]interface{}) error {
	channelAttrs := dto.ChannelAttributes{
		Description: attrs["description"].(string),
		IsPrivate:   attrs["is_private"].(bool),
	}
	return validator.Validate(channelAttrs)
}

func (c *ChannelCreator) CreateEntity(chatType, title string, userID uuid.UUID, attrs map[string]interface{}) (entity.Chat, error) {
	filteredAttrs := c.FilterAttributes(attrs, []string{"description", "is_private"})
	return entity.NewChat(chatType, title, userID, filteredAttrs), nil
}

//методы для диологов

type DiologCreator struct {
	BaseChatCreator
}

func (c *DiologCreator) ValidateAttributes(attrs map[string]interface{}) error {
	diologAttrs := dto.DiologAttributes{
		User1ID: attrs["user1_id"].(uuid.UUID),
		User2ID: attrs["user2_id"].(uuid.UUID),
	}
	return validator.Validate(diologAttrs)
}

func (c *DiologCreator) CreateEntity(chatType, title string, userID uuid.UUID, attrs map[string]interface{}) (entity.Chat, error) {
	filteredAttrs := c.FilterAttributes(attrs, []string{"user1_id", "user2_id"})
	return entity.NewChat(chatType, title, userID, filteredAttrs), nil
}

//методы для групп

type GroupCreator struct {
	BaseChatCreator
}

func (c *GroupCreator) ValidateAttributes(attrs map[string]interface{}) error {
	groupAttrs := dto.GroupAttributes{
		IsPublic: attrs["is_public"].(bool),
	}
	return validator.Validate(groupAttrs)
}

func (c *GroupCreator) CreateEntity(chatType, title string, userID uuid.UUID, attrs map[string]interface{}) (entity.Chat, error) {
	filteredAttrs := c.FilterAttributes(attrs, []string{"is_public"})
	return entity.NewChat(chatType, title, userID, filteredAttrs), nil
}
