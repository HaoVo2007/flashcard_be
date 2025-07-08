package topics

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicRepository interface {
	CreateTopic(c context.Context, topic *Topic) error
	GetAllTopics(c context.Context) ([]*Topic, error)
	GetTopicByID(c context.Context, id primitive.ObjectID) (*Topic, error)
	GetTopicsByUserID(c context.Context, userID primitive.ObjectID) ([]*Topic, error)
	UpdateTopic(c context.Context, id primitive.ObjectID, topic *Topic) error
	DeleteTopic(c context.Context, id primitive.ObjectID) error
}

type topicRepository struct {
	collection *mongo.Collection
}

func NewTopicRepository(collection *mongo.Collection) TopicRepository {
	return &topicRepository{collection: collection}
}

func (r *topicRepository) CreateTopic(c context.Context, topic *Topic) error {
	_, err := r.collection.InsertOne(c, topic)
	if err != nil {
		return err
	}
	return nil
}

func (r *topicRepository) GetAllTopics(c context.Context) ([]*Topic, error) {

	var topics []*Topic
	
	cursor, err := r.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	for cursor.Next(c) {
		var topic Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, &topic)
	}

	return topics, nil

}

func (r *topicRepository) GetTopicByID(c context.Context, id primitive.ObjectID) (*Topic, error) {

	var topic Topic
	
	err := r.collection.FindOne(c, bson.M{"_id": id}).Decode(&topic)
	if err != nil {
		return nil, err
	}
	
	return &topic, nil

}

func (r *topicRepository) GetTopicsByUserID(c context.Context, userID primitive.ObjectID) ([]*Topic, error) {

	var topics []*Topic
	
	cursor, err := r.collection.Find(c, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	for cursor.Next(c) {
		var topic Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, &topic)
	}

	return topics, nil
}
func (r *topicRepository) UpdateTopic(c context.Context, id primitive.ObjectID, topic *Topic) error {

	filter := bson.M{"_id": id}
	update := bson.M{"$set": topic}

	_, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *topicRepository) DeleteTopic(c context.Context, id primitive.ObjectID) error {

	_, err := r.collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil

}

