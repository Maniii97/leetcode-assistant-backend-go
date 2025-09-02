package handlers

import (
	"lc-assistant-backend/internal/services"
	"lc-assistant-backend/pkg/request"
	"lc-assistant-backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SolutionHandler struct {
	genaiService service.GenAIService
}

func NewSolutionHandler(genaiService service.GenAIService) *SolutionHandler {
	return &SolutionHandler{
		genaiService: genaiService,
	}
}

func (h *SolutionHandler) AnalyseSolution(c *gin.Context) {
	var req request.AnalysisRequest

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
	feedback, err := h.genaiService.AnalyseSolution(req.Question, req.Solution)
	if err != nil {
		resp := response.Error{
			Success: false,
			Message: "Failed to get feedback, Retry",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.AnalysisSuccess{
		Success:  true,
		Feedback: feedback,
	}

	c.JSON(http.StatusOK, resp)
}
