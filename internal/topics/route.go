package topics

import (
	"flashcard/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *TopicHandler) {

	topicGroup := r.Group("/api/v1/topic")
	{
		topicGroup.POST("", middleware.JWTAuthMiddleware(), handler.CreateTopic)
		topicGroup.GET("", middleware.JWTAuthMiddleware(), handler.GetAllTopics)
		topicGroup.GET("/:topic_id", middleware.JWTAuthMiddleware(), handler.GetTopicByID)
		topicGroup.GET("/user", middleware.JWTAuthMiddleware(), handler.GetAllTopicsByUser)
		topicGroup.PUT("/:topic_id", middleware.JWTAuthMiddleware(), handler.UpdateTopic)
		topicGroup.DELETE("/:topic_id", middleware.JWTAuthMiddleware(), handler.DeleteTopic)
	}

}