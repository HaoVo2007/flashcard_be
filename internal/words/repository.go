package words

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordRepository interface{
	CreateWord(c context.Context, word *Word) error
	GetAllWords(c context.Context) ([]*Word, error)
	GetWordByID(c context.Context, id primitive.ObjectID) (*Word, error)
	GetWordsByTopicID(c context.Context, id primitive.ObjectID) ([]*Word, error)
}

type wordRepository struct {
	collection *mongo.Collection
}

func NewWordRepository(collection *mongo.Collection) WordRepository {
	return &wordRepository{collection: collection}
}

func (r *wordRepository) CreateWord(c context.Context, word *Word) error {
	_, err := r.collection.InsertOne(c, word)
	if err != nil {
		return err
	}
	return nil
}

func (r *wordRepository) GetAllWords(c context.Context) ([]*Word, error) {

	var words []*Word
	
	cursor, err := r.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	for cursor.Next(c) {
		var word Word
		if err := cursor.Decode(&word); err != nil {
			return nil, err
		}
		words = append(words, &word)
	}
	
	return words, nil

}

func (r *wordRepository) GetWordByID(c context.Context, id primitive.ObjectID) (*Word, error) {
	
	filter := bson.M{"_id": id}
	
	var word Word
	
	err := r.collection.FindOne(c, filter).Decode(&word)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	
	return &word, nil

}

func (r *wordRepository) GetWordsByTopicID(c context.Context, id primitive.ObjectID) ([]*Word, error) {
	
	filter := bson.M{"topic_id": id}
	
	var words []*Word
	
	cursor, err := r.collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	for cursor.Next(c) {
		var word Word
		if err := cursor.Decode(&word); err != nil {
			return nil, err
		}
		words = append(words, &word)
	}
	
	return words, nil

}