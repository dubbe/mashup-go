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
	d := make(chan Description)
	e := make(chan error)
	go wikipedia.Get(d, e)
	desc := <-d
	err := <-e
	assert.Equal(t, err, nil)
	assert.Contains(t, desc, "The Temptations")
}

func TestGetEbbaGron(t *testing.T) {
	wikipedia := NewWikipedia(&http.Client{})
	wikipedia.SetRelation(artist.Relation{
		Type: "wikipedia",
		URL: artist.URL{
			Resource: "https://en.wikipedia.org/wiki/Ebba_Gr%C3%B6n",
		},
	})
	d := make(chan Description)
	e := make(chan error)
	go wikipedia.Get(d, e)
	desc := <-d
	err := <-e
	assert.Equal(t, nil, err)
	assert.Contains(t, desc, "Ebba Grön")
}

func TestGetUnknown(t *testing.T) {
	wikipedia := NewWikipedia(&http.Client{})
	wikipedia.SetRelation(artist.Relation{
		Type: "wikipedia",
		URL: artist.URL{
			Resource: "https://en.wikipedia.org/wiki/uknown2222",
		},
	})
	d := make(chan Description)
	e := make(chan error)
	go wikipedia.Get(d, e)
	desc := <-d
	err := <-e
	assert.Equal(t, nil, err)
	assert.Contains(t, desc, "")
}
