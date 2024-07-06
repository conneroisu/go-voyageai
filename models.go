package voyageai

// modelName is the name of a model
type modelName string

const (
	// VoyageLarge2Instruct is the model for large language models
	VoyageLarge2Instruct = modelName("voyage-large-2-instruct")
	// VoyageFinace2 is the model for finance
	VoyageFinace2 = modelName("voyage-finance-2")
	// VoyageMultilingual2 is the model for multilingual
	VoyageMultilingual2 = modelName("voyage-multilingual-2")
	// VoyageLaw2 is the model for law
	VoyageLaw2 = modelName("voyage-law-2")
	// VoyageCode2 is the model for code
	VoyageCode2 = modelName("voyage-code-2")
	// VoyageLarge2 is the model for large language models
	VoyageLarge2 = modelName("voyage-large-2")
	// Voyage2 is the model for large language models
	Voyage2 = modelName("voyage-2")
)

const defaultBaseURL = "https://api.voyage.ai/v1"

// EmbeddingsRequest is the request to the voyage ai api
type EmbeddingsRequest struct {
	// Object is the type of the request, it is always "list"
	Model string `json:"model"`
	// Input is an array of strings that contain the input to be embedded.
	Input []string `json:"input"`
}

// EmbeddingsResponse is the response from the voyage ai api
type EmbeddingsResponse struct {
	// Object is the type of the response, it is always "list"
	Object string `json:"object"`
	// Data is an array of embedding objects that contain the embedding and the index of the input.
	Data []struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	// Usage is the usage of the model for the request.
	Usage struct {
		// TotalTokens is the total number of tokens used by the model for the request.
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

// RerankRequest is the request to the voyage ai api
type RerankRequest struct {
	Query      string   `json:"query"`
	Documents  []string `json:"documents"`
	Model      string   `json:"model"`
	TopK       int      `json:"top_k,omitempty"`
	Truncation bool     `json:"truncation,omitempty"`
}

// RerankResponse is the response from the voyage ai api
type RerankResponse struct {
	Object string `json:"object"`
	Data   []struct {
		RelevanceScore float64 `json:"relevance_score"`
		Index          int     `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}
