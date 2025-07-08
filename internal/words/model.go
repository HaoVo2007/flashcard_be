package words

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Word struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	TopicID    primitive.ObjectID `json:"topic_id" bson:"topic_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Word       string             `json:"word" bson:"word"`
	Definition string             `json:"definition" bson:"definition"`
	Example    *string             `json:"example" bson:"example"`
	WordType   string             `json:"word_type" bson:"word_type"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}