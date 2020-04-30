package description

import (
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/stretchr/testify/assert"
)

func TestWikidataGetNirvana(t *testing.T) {
	wikipedia := NewWikipedia(&http.Client{})

	wikidata := NewWikidata(&http.Client{}, wikipedia)

	wikidata.SetRelation(artist.Relation{
		Type: "wikidata",
		URL: artist.URL{
			Resource: "https://www.wikidata.org/wiki/Q11649",
		},
	})
	description, err := wikidata.Get()
	assert.Equal(t, nil, err)
	assert.Contains(t, description.Description, "Nirvana")
}
