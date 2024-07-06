package voyageai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is a client for voyage ai
type Client struct {
	apiKey string
	base   string
}

// NewClient creates a new client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		base:   defaultBaseURL,
	}
}

// Embeddings returns embeddings for a given text
func (c *Client) Embeddings(req EmbeddingsRequest) (EmbeddingsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return EmbeddingsResponse{}, err
	}
	request := http.Request{
		Method: "POST",
		URL:    &url.URL{Path: c.base + "/embeddings"},
		Header: http.Header{
			"Authorization": []string{"Bearer " + c.apiKey},
			"content-type":  []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	response, err := http.DefaultClient.Do(&request)
	if err != nil {
		return EmbeddingsResponse{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return EmbeddingsResponse{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	var res EmbeddingsResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return EmbeddingsResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return res, nil
}

// Rerank returns reranked results of a given query
func (c *Client) Rerank(req RerankRequest) (RerankResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return RerankResponse{}, err
	}
	request := http.Request{
		Method: "POST",
		URL:    &url.URL{Path: c.base + "/rerank"},
		Header: http.Header{
			"Authorization": []string{"Bearer " + c.apiKey},
			"content-type":  []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	response, err := http.DefaultClient.Do(&request)
	if err != nil {
		return RerankResponse{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return RerankResponse{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	var res RerankResponse
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return RerankResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return res, nil
}
