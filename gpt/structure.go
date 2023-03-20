package gpt

type completionsRequest struct {
	Model       string      `json:"model"`
	Prompt      string      `json:"prompt"`
	MaxTokens   int         `json:"max_tokens"`
	Temperature int         `json:"temperature"`
	TopP        int         `json:"top_p,omitempty"`
	N           int         `json:"n,omitempty"`
	Stream      bool        `json:"stream,omitempty"`
	Logprobs    interface{} `json:"logprobs,omitempty"`
	Stop        string      `json:"stop,omitempty"`
}

type completionsResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
