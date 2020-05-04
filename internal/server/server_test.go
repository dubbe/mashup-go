package server

import (
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/internal/client/coverart"
	"github.com/dubbe/mashup-go/internal/client/description"
	"github.com/stretchr/testify/assert"
)

var s *Server

func init() {
	wikipedia := description.NewWikipedia(&http.Client{})

	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	httpClient := &http.Client{}

	s = NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
}

func TestGetArtist(t *testing.T) {
	artist, err := s.getArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Nirvana", artist.Name)
}

func BenchmarkArtist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.getArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	}
}
