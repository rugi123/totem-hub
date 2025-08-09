package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	MemberID primitive.ObjectID `bson:"memberID"`
	Text     string             `bson:"text,omitempty"`
	SentAt   time.Time          `bson:"sentAt"`
	EditedAt *time.Time         `bson:"editedAt,omitempty"`

	//вложеный контент
	Content map[string]interface{} `bson:"content,omitempty"`
}
