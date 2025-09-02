package handlers

import (
	"lc-assistant-backend/internal/services"
	"lc-assistant-backend/pkg/request"
	"lc-assistant-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	genaiService service.GenAIService
}

func NewReviewHandler(genaiService service.GenAIService) *ReviewHandler {
	return &ReviewHandler{
		genaiService: genaiService,
	}
}

func (h *ReviewHandler) ReviewCode(c *gin.Context) {
	var req request.ReviewRequest

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

	// Use the genai service to get response
	feedback, err := h.genaiService.ReviewCode(req.Code, req.Question)
	if err != nil {
		resp := response.Error{
			Success: false,
			Message: "Failed to get feedback, Retry",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.ReviewSuccess{
		Success: true,
		Review:  feedback,
	}

	c.JSON(http.StatusOK, resp)
}
