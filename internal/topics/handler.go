package topics

import (
	"flashcard/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	TopicService TopicService
}

func NewTopicHandler(topicService TopicService) *TopicHandler {
	return &TopicHandler{TopicService: topicService}
}

func (h *TopicHandler) CreateTopic(c *gin.Context) {

	var req CreateTopicRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		helper.SendError(c, http.StatusUnauthorized, nil, helper.ErrInvalidOperation)
		return
	}
	err := h.TopicService.CreateTopic(c, &req, userID.(string))

	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "success", nil)
}

func (h *TopicHandler) GetAllTopics(c *gin.Context) {

	topics, err := h.TopicService.GetAllTopics(c)

	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", topics)
}

func (h *TopicHandler) GetTopicByID(c *gin.Context) {
	
	id := c.Param("topic_id")

	topic, err := h.TopicService.GetTopicByID(c, id)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", topic)
	
}

func (h *TopicHandler) GetAllTopicsByUser(c *gin.Context) {

	userID, ok := c.Get("user_id")
	if !ok {
		helper.SendError(c, http.StatusUnauthorized, nil, helper.ErrInvalidOperation)
		return
	}

	topics, err := h.TopicService.GetTopicsByUserID(c, userID.(string))
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", topics)

}
func (h *TopicHandler) UpdateTopic(c *gin.Context) {

	id := c.Param("topic_id")

	var req UpdateTopicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	err := h.TopicService.UpdateTopic(c, id, &req)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", nil)
}

func (h *TopicHandler) DeleteTopic(c *gin.Context) {

	id := c.Param("topic_id")

	err := h.TopicService.DeleteTopic(c, id)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", nil)

}