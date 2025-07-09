package topics

import (
	"context"
	"flashcard/internal/words"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopicService interface {
	CreateTopic(c context.Context, req *CreateTopicRequest, userID string) error
	GetAllTopics(c context.Context) ([]*Topic, error)
	GetTopicByID(c context.Context, id string) (*TopicResponse, error)
	GetTopicsByUserID(c context.Context, userID string) ([]*Topic, error)
	UpdateTopic(c context.Context, id string, req *UpdateTopicRequest) error
	DeleteTopic(c context.Context, id string) error
}

type topicService struct {
	topicRepository TopicRepository
	wordService words.WordService
}

func NewTopicService(topicRepository TopicRepository, wordService words.WordService) TopicService {
	return &topicService{
		topicRepository: topicRepository,
		wordService:     wordService,
	}
}

func (s *topicService) CreateTopic(c context.Context, req *CreateTopicRequest, userID string) error {
	
	if req.Color == "" {
		req.Color = "#000000"
	}

	if req.TopicName == "" {
		return fmt.Errorf("topic name is required")
	}

	if userID == "" {
		return fmt.Errorf("user id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	topic := &Topic{
		ID :              primitive.NewObjectID(),
		TopicName:        req.TopicName,
		TopicDescription: &req.TopicDescription,
		Color:            req.Color,
		UserID:           objectID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return s.topicRepository.CreateTopic(c, topic)

}

func (s *topicService) GetAllTopics(c context.Context) ([]*Topic, error) {
	return s.topicRepository.GetAllTopics(c)
}

func (s *topicService) GetTopicByID(c context.Context, id string) (*TopicResponse, error) {
	
	if id == "" {
		return nil, fmt.Errorf("topic id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	topic, err := s.topicRepository.GetTopicByID(c, objectID)
	if err != nil {
		return nil, err
	}

	works, err := s.wordService.GetWordsByTopicID(c, topic.ID.Hex())
	if err != nil {
		return nil, err
	}

	res := &TopicResponse{
		ID:               topic.ID,
		TopicName:        topic.TopicName,
		TopicDescription: topic.TopicDescription,
		Color:            topic.Color,
		UserID:           topic.UserID,
		CreatedAt:        topic.CreatedAt,
		UpdatedAt:        topic.UpdatedAt,
		WordCount:        len(works),
	}

	return res, nil
}

func (s *topicService) GetTopicsByUserID(c context.Context, userID string) ([]*Topic, error) {
	
	if userID == "" {
		return nil, fmt.Errorf("user id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return s.topicRepository.GetTopicsByUserID(c, objectID)
	
}

func (s *topicService) UpdateTopic(c context.Context, id string, req *UpdateTopicRequest) error {
	
	if id == "" {
		return fmt.Errorf("topic id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	topic, er := s.topicRepository.GetTopicByID(c, objectID)
	if er != nil {
		return er
	}

	if req.Color != nil {
		topic.Color = *req.Color
	}

	if req.TopicName != nil {
		topic.TopicName = *req.TopicName
	}

	if req.TopicDescription != nil {
		topic.TopicDescription = req.TopicDescription
	}

	return s.topicRepository.UpdateTopic(c, objectID, topic)

}

func (s *topicService) DeleteTopic(c context.Context, id string) error {
	
	if id == "" {
		return fmt.Errorf("topic id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	words, err := s.wordService.GetWordsByTopicID(c, id)
	if err != nil {
		return err
	}

	for _, word := range words {
		err := s.wordService.DeleteWord(c, word.ID.Hex())
		if err != nil {
			return err
		}
	}

	return s.topicRepository.DeleteTopic(c, objectID)

}