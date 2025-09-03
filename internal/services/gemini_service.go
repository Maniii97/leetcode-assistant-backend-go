package service

import (
	"context"
	"fmt"
	"lc-assistant-backend/internal/constants"
	"os"
	"sync"

	"google.golang.org/genai"
)

// GeminiConfig holds configuration for the Gemini service
type GeminiConfig struct {
	SystemInstruction string
	ModelName         string
	APIKey            string
}

type geminiService struct {
	ctx    context.Context
	client *genai.Client
	chat   *genai.Chat
	config *GeminiConfig
	mutex  sync.RWMutex
}

// NewGeminiService creates a new instance of the Gemini service
func NewGeminiService(config *GeminiConfig) (GenAIService, error) {
	if config == nil {
		config = &GeminiConfig{
			SystemInstruction: constants.SystemInstruction,
			ModelName:         constants.Model,
			APIKey:            os.Getenv("GEMINI_API_KEY"),
		}
	}

	if config.APIKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is required")
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: config.APIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	// Create chat configuration with system instruction
	chatConfig := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(config.SystemInstruction, genai.RoleUser),
	}

	chat, err := client.Chats.Create(ctx, config.ModelName, chatConfig, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat session: %w", err)
	}

	service := &geminiService{
		ctx:    ctx,
		client: client,
		chat:   chat,
		config: config,
	}

	return service, nil
}

func (g *geminiService) sendMessage(message string) (string, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	response, err := g.chat.SendMessage(g.ctx, genai.Part{Text: message})
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	return response.Text(), nil
}

func (g *geminiService) GetChatResponse(prompt string) (string, error) {
	return g.sendMessage(prompt)
}

func (g *geminiService) AnalyseSolution(question string, solution string) (string, error) {
	prompt := fmt.Sprintf(constants.SolutionAnalysisPrompt, question, solution)
	return g.sendMessage(prompt)
}

func (g *geminiService) GetHints(question string) (string, error) {
	prompt := fmt.Sprintf(constants.GetHintsPrompt, question)
	return g.sendMessage(prompt)
}

func (g *geminiService) ReviewCode(code string, question string) (string, error) {
	prompt := fmt.Sprintf(constants.CodeReviewPrompt, code, question)
	return g.sendMessage(prompt)
}

func (g *geminiService) ClearChatHistory() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	chatConfig := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(g.config.SystemInstruction, genai.RoleUser),
	}

	chat, err := g.client.Chats.Create(g.ctx, g.config.ModelName, chatConfig, nil)
	if err != nil {
		return fmt.Errorf("failed to create new chat session: %w", err)
	}

	g.chat = chat
	return nil
}

func (g *geminiService) GetChatHistory() []*genai.Content {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.chat.History(false)
}

// Close cleans up the service resources
func (g *geminiService) Close() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// The genai client doesn't expose a Close method,
	// but we can clear the references to help with garbage collection
	g.chat = nil
	g.client = nil
	return nil
}
