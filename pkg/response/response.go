package response

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type ChatSuccess struct {
	Success bool   `json:"success"`
	Answer  string `json:"answer"`
}

type HintSuccess struct {
	Success bool   `json:"success"`
	Hints   string `json:"hints"`
}

type ReviewSuccess struct {
	Success bool   `json:"success"`
	Review  string `json:"review"`
}

type AnalysisSuccess struct {
	Success  bool   `json:"success"`
	Feedback string `json:"feedback"`
}
