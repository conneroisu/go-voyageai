package voyageai

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEmbeddingRequest tests the embedding request.
//
// it ensures that the request serializes and deserializes correctly.
func TestEmbeddingRequest(t *testing.T) {
	a := assert.New(t)
	// create a new embedding request and make sure that it serializes correctly
	req := EmbeddingsRequest{
		Model: string(VoyageLarge2Instruct),
		Input: []string{"Hello, world!"},
	}
	a.Equal(req.Model, string(VoyageLarge2Instruct))
	a.Equal(req.Input, []string{"Hello, world!"})
	b, err := json.Marshal(req)
	a.NoError(err)
	c := EmbeddingsRequest{}
	err = json.Unmarshal(b, &c)
	a.NoError(err)
	a.Equal(c.Model, string(VoyageLarge2Instruct))
	a.Equal(c.Input, []string{"Hello, world!"})
}
