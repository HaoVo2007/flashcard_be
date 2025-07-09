package words

import (
	"flashcard/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	WordService WordService
}

func NewWordHandler(wordService WordService) *WordHandler {
	return &WordHandler{WordService: wordService}
}

func (h *WordHandler) CreateWord(c *gin.Context) {

	userID, ok := c.Get("user_id")
	if !ok {
		helper.SendError(c, http.StatusUnauthorized, nil, helper.ErrInvalidOperation)
		return
	}

	var req CreateWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	err := h.WordService.CreateWord(c, &req, userID.(string))
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "success", nil)
}

func (h *WordHandler) GetAllWords(c *gin.Context) {

	words, err := h.WordService.GetAllWords(c)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", words)

}

func (h *WordHandler) GetWordByID(c *gin.Context) {

	id := c.Param("word_id")

	word, err := h.WordService.GetWordByID(c, id)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", word)

}

func (h *WordHandler) GetAllWordsByTopicID(c *gin.Context) {

	id := c.Param("topic_id")

	words, err := h.WordService.GetWordsByTopicID(c, id)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", words)
	
}

func (h *WordHandler) UpdateWord(c *gin.Context) {

	id := c.Param("word_id")

	var req UpdateWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	err := h.WordService.UpdateWord(c, id, &req)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", nil)

}

func (h *WordHandler) DeleteWord(c *gin.Context) {

	id := c.Param("word_id")

	err := h.WordService.DeleteWord(c, id)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "success", nil)
	
}