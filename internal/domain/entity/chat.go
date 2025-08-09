package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Type      string             `bson:"type"`
	Title     string             `bson:"title"`
	CreatedBy primitive.ObjectID `bson:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt"`

	//динамичекские настройки чата
	Settings map[string]interface{} `bson:"settings,omitempty"`
}

type ChatMember struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"userID,omitempty"`
	ChatID   primitive.ObjectID `bson:"chatID,omitempty"`
	JoinedAt time.Time          `bson:"joinedAt"`
	Role     string             `bson:"role"`

	//динамичные права и настройки
	Permissions map[string]bool        `bson:"permissions,omitempty"`
	CustomData  map[string]interface{} `bson:"customData,omitempty"`
}
