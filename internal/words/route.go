package words

import (
	"flashcard/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *WordHandler) {

	wordGroup := r.Group("/api/v1/word")
	{
		wordGroup.POST("", middleware.JWTAuthMiddleware(), handler.CreateWord)
		wordGroup.GET("", middleware.JWTAuthMiddleware(), handler.GetAllWords)
		wordGroup.GET("/:word_id", middleware.JWTAuthMiddleware(), handler.GetWordByID)
		wordGroup.GET("/topic/:topic_id", middleware.JWTAuthMiddleware(), handler.GetAllWordsByTopicID)
		wordGroup.PUT("/:word_id", middleware.JWTAuthMiddleware(), handler.UpdateWord)
		wordGroup.DELETE("/:word_id", middleware.JWTAuthMiddleware(), handler.DeleteWord)
	}

}