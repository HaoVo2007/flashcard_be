package words

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordRepository interface{
	CreateWord(c context.Context, word *Word) error
	GetAllWords(c context.Context, req *SearchWordRequest) ([]*Word, error)
	GetWordByID(c context.Context, id primitive.ObjectID) (*Word, error)
	GetWordsByTopicID(c context.Context, id primitive.ObjectID, req *SearchWordRequest) ([]*Word, error)
	UpdateWord(c context.Context, id primitive.ObjectID, word *Word) error
	DeleteWord(c context.Context, id primitive.ObjectID) error
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

func (r *wordRepository) GetAllWords(c context.Context, req *SearchWordRequest) ([]*Word, error) {

	var words []*Word

	filter := bson.M{}
	if req.Istrue != nil {
		filter["is_true"] = *req.Istrue
	}
	if req.TopicID != nil {
		filter["topic_id"] = *req.TopicID
	}
	if req.Word != nil {
		filter["word"] = *req.Word
	}

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

func (r *wordRepository) GetWordsByTopicID(c context.Context, id primitive.ObjectID, req *SearchWordRequest) ([]*Word, error) {
	
	filter := bson.M{"topic_id": id}
	
	if req.Word != nil {
		filter["word"] = *req.Word
	}

	if req.Istrue != nil {
		filter["is_true"] = *req.Istrue
	}
	
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

func (r *wordRepository) UpdateWord(c context.Context, id primitive.ObjectID, word *Word) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": word}
	_, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *wordRepository) DeleteWord(c context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
}