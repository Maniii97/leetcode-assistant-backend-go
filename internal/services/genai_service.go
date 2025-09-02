package service

type GenAIService interface {
	GetHints(question string) (string, error)
	GetChatResponse(prompt string) (string, error)
	AnalyseSolution(question string, solution string) (string, error)
	ReviewCode(code string, question string) (string, error)
	Close() error
}
