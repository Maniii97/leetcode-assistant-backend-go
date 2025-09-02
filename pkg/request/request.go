package request

import (
	"errors"
	"strings"
)

type ChatRequest struct {
	Question string `json:"question"`
}

type HintRequest struct {
	Question string `json:"question"`
}

type ReviewRequest struct {
	Code     string `json:"code"`
	Question string `json:"question"`
}

type AnalysisRequest struct {
	Question string `json:"question"`
	Solution string `json:"solution"`
}

func (c *AnalysisRequest) Validate() error {
	if strings.TrimSpace(c.Question) == "" {
		return errors.New("question is required")
	}
	if strings.TrimSpace(c.Solution) == "" {
		return errors.New("solution is required")
	}
	return nil
}

func (c *ChatRequest) Validate() error {
	if strings.TrimSpace(c.Question) == "" {
		return errors.New("question is required")
	}
	return nil
}

func (h *HintRequest) Validate() error {
	if strings.TrimSpace(h.Question) == "" {
		return errors.New("question is required")
	}
	return nil
}

func (r *ReviewRequest) Validate() error {
	if strings.TrimSpace(r.Code) == "" {
		return errors.New("code is required")
	}
	return nil
}
