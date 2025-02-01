package openAI

type RequestPayload struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type ResponsePayload struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	ServiceTier       string   `json:"service_tier"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

// Choice Estructura para las elecciones en la respuesta
type Choice struct {
	Index        int         `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"` // null o una estructura, si se espera datos
	FinishReason string      `json:"finish_reason"`
}

// Message Estructura para los mensajes dentro de las elecciones
type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"` // null o una estructura, si se espera datos
}

// Usage Estructura para el uso detallado de tokens y otros detalles
type Usage struct {
	PromptTokens            int                    `json:"prompt_tokens"`
	CompletionTokens        int                    `json:"completion_tokens"`
	TotalTokens             int                    `json:"total_tokens"`
	PromptTokensDetails     TokenDetails           `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokenDetails `json:"completion_tokens_details"`
}

// TokenDetails Detalles sobre los tokens del prompt
type TokenDetails struct {
	CachedTokens int `json:"cached_tokens"`
	AudioTokens  int `json:"audio_tokens"`
}

// CompletionTokenDetails Detalles sobre los tokens de la completación
type CompletionTokenDetails struct {
	ReasoningTokens          int `json:"reasoning_tokens"`
	AudioTokens              int `json:"audio_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}

type ChatRequest struct {
	Model     string           `json:"model"`
	Store     bool             `json:"store"`
	Messages  []MessageRequest `json:"messages"`
	MaxTokens int              `json:"max_tokens,omitempty"`
}

// MessageRequest representa cada mensaje en el array de "messages" en la solicitud
type MessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
