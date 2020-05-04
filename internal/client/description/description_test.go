package description

import (
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/stretchr/testify/assert"
)

var descriptionFactory *DescriptionFactory

func init() {
	/* load test data */
	wikipedia := NewWikipedia(&http.Client{})

	wikidata := NewWikidata(&http.Client{}, wikipedia)

	descriptionFactory = CreateDescriptionFactory(wikipedia, wikidata)
}

func TestDescriptionWikidata(t *testing.T) {

	client, err := descriptionFactory.NewDescriptionClient(artist.Relation{
		Type: "wikidata",
		URL: artist.URL{
			Resource: "https://www.wikidata.org/wiki/Q11649",
		},
	})

	assert.Equal(t, nil, err)

	d := make(chan Description)
	e := make(chan error)
	go client.Get(d, e)
	desc := <-d
	err = <-e

	assert.Equal(t, nil, err)
	assert.Contains(t, desc, "Nirvana")
}

func TestDescriptionWikipedia(t *testing.T) {

	client, err := descriptionFactory.NewDescriptionClient(artist.Relation{
		Type: "wikipedia",
		URL: artist.URL{
			Resource: "https://en.wikipedia.org/wiki/The_Temptations",
		},
	})
	assert.Equal(t, nil, err)

	d := make(chan Description)
	e := make(chan error)
	go client.Get(d, e)
	desc := <-d
	err = <-e

	assert.Equal(t, nil, err)
	assert.Contains(t, desc, "The Temptations")
}
