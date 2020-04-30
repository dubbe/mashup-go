package description

import (
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/stretchr/testify/assert"
)

func TestGetTheTemptations(t *testing.T) {
	wikipedia := NewWikipedia(&http.Client{})
	wikipedia.SetRelation(artist.Relation{
		Type: "wikipedia",
		URL: artist.URL{
			Resource: "https://en.wikipedia.org/wiki/The_Temptations",
		},
	})
	description, err := wikipedia.Get()
	assert.Equal(t, nil, err)
	assert.Contains(t, description.Description, "The Temptations")
}
