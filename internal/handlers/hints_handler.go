package handlers

import (
	"lc-assistant-backend/internal/services"
	"lc-assistant-backend/pkg/request"
	"lc-assistant-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HintsHandler struct {
	genaiService service.GenAIService
}

func NewHintsHandler(genaiService service.GenAIService) *HintsHandler {
	return &HintsHandler{
		genaiService: genaiService,
	}
}

func (h *HintsHandler) GetHints(c *gin.Context) {
	var req request.HintRequest

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
	answer, err := h.genaiService.GetHints(req.Question)
	if err != nil {
		resp := response.Error{
			Success: false,
			Message: "Failed to get answer",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.HintSuccess{
		Success: true,
		Hints:   answer,
	}

	c.JSON(http.StatusOK, resp)
}
