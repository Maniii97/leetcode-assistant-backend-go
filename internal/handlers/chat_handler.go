package handlers

import (
	"lc-assistant-backend/internal/services"
	"lc-assistant-backend/pkg/request"
	"lc-assistant-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	genaiService service.GenAIService
}

func NewChatHandler(genaiService service.GenAIService) *ChatHandler {
	return &ChatHandler{
		genaiService: genaiService,
	}
}

func (h *ChatHandler) GetAnswer(c *gin.Context) {
	var req request.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp := response.Error{
			Success: false,
			Message: "Failed to bind request",
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		resp := response.Error{
			Success: false,
			Message: "Invalid request",
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Use the genai service to get chat response
	answer, err := h.genaiService.GetChatResponse(req.Question)
	if err != nil {
		resp := response.Error{
			Success: false,
			Message: "Failed to get answer",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.ChatSuccess{
		Success: true,
		Answer:  answer,
	}

	c.JSON(http.StatusOK, resp)
}
