package server

import (
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/internal/client/coverart"
	"github.com/dubbe/mashup-go/internal/client/description"
	"github.com/stretchr/testify/assert"
)

func TestGetArtist(t *testing.T) {

	wikipedia := description.NewWikipedia(&http.Client{})

	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	httpClient := &http.Client{}

	server := NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
	//server.Run()

	artist := server.getArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	assert.Equal(t, artist.Name, "Nirvana")
}

func TestGetArtistAsync(t *testing.T) {

	wikipedia := description.NewWikipedia(&http.Client{})

	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	httpClient := &http.Client{}

	server := NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
	//server.Run()

	artist := server.getArtistAsync("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	assert.Equal(t, artist.Name, "Nirvana")
}

func BenchmarkArtist(b *testing.B) {
	wikipedia := description.NewWikipedia(&http.Client{})

	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	httpClient := &http.Client{}

	server := NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
	//server.Run()
	for i := 0; i < b.N; i++ {
		server.getArtist("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	}
	// assert.Contains(t, description.Description, "The Temptations")
}

func BenchmarkArtistAsync(b *testing.B) {
	wikipedia := description.NewWikipedia(&http.Client{})

	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	httpClient := &http.Client{}

	server := NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
	//server.Run()
	for i := 0; i < b.N; i++ {
		server.getArtistAsync("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	}
	// assert.Contains(t, description.Description, "The Temptations")
}
