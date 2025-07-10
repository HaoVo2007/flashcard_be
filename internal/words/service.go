package words

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WordService interface {
	CreateWord(c context.Context, req *CreateWordRequest, userID string) error
	GetAllWords(c context.Context, req *SearchWordRequest) ([]*Word, error)
	GetWordByID(c context.Context, id string) (*Word, error)
	GetWordsByTopicID(c context.Context, id string, req *SearchWordRequest) ([]*Word, error)
	UpdateWord(c context.Context, id string, req *UpdateWordRequest) error
	DeleteWord(c context.Context, id string) error
}

type wordService struct {
	wordRepository WordRepository
}

func NewWordService(wordRepository WordRepository) WordService {
	return &wordService{wordRepository: wordRepository}
}

func (s *wordService) CreateWord(c context.Context, req *CreateWordRequest, userID string) error {

	if req.TopicID == "" {
		return fmt.Errorf("topic id is required")
	}

	if req.Word == "" {
		return fmt.Errorf("word is required")
	}

	if req.Definition == "" {
		return fmt.Errorf("definition is required")
	}

	if userID == "" {
		return fmt.Errorf("user id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	objectTopicID, err := primitive.ObjectIDFromHex(req.TopicID)
	if err != nil {
		return err
	}

	word := &Word{
		ID:         primitive.NewObjectID(),
		TopicID:    objectTopicID,
		UserID:     objectID,
		Word:       req.Word,
		Definition: req.Definition,
		Example:    req.Example,
		IsTrue:     false,
		WordType:   req.WordType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.wordRepository.CreateWord(c, word)

}

func (s *wordService) GetAllWords(c context.Context, req *SearchWordRequest) ([]*Word, error) {
	return s.wordRepository.GetAllWords(c, req)
}

func (s *wordService) GetWordByID(c context.Context, id string) (*Word, error) {

	if id == "" {
		return nil, fmt.Errorf("word id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return s.wordRepository.GetWordByID(c, objectID)

}

func (s *wordService) GetWordsByTopicID(c context.Context, id string, req *SearchWordRequest) ([]*Word, error) {

	if id == "" {
		return nil, fmt.Errorf("topic id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return s.wordRepository.GetWordsByTopicID(c, objectID, req)

}

func (s *wordService) UpdateWord(c context.Context, id string, req *UpdateWordRequest) error {

	if id == "" {
		return fmt.Errorf("word id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	word, err := s.wordRepository.GetWordByID(c, objectID)
	if err != nil {
		return err
	}

	if req.Word != nil {
		word.Word = *req.Word
	}

	if req.Definition != nil {
		word.Definition = *req.Definition
	}

	if req.Example != nil {
		word.Example = req.Example
	}

	if req.WordType != nil {
		word.WordType = *req.WordType
	}

	if req.IsTrue != nil {
		word.IsTrue = *req.IsTrue
	}

	return s.wordRepository.UpdateWord(c, objectID, word)
}

func (s *wordService) DeleteWord(c context.Context, id string) error {

	if id == "" {
		return fmt.Errorf("word id is required")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.wordRepository.DeleteWord(c, objectID)

}
