package topics

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topic struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TopicName        string             `json:"name" bson:"name"`
	TopicDescription *string             `json:"description" bson:"description"`
	Color            string             `json:"color" bson:"color"`
	UserID           primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}
