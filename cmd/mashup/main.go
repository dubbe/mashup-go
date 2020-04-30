package main

import (
	"net/http"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/internal/client/coverart"
	"github.com/dubbe/mashup-go/internal/client/description"
	"github.com/dubbe/mashup-go/internal/server"
)

func main() {
	httpClient := &http.Client{}

	wikipedia := description.NewWikipedia(&http.Client{})
	wikidata := description.NewWikidata(&http.Client{}, wikipedia)

	s := server.NewServer(artist.NewMusicbrainz(httpClient), description.CreateDescriptionFactory(wikipedia, wikidata), coverart.NewCoverartArchive(httpClient))
	s.Run()
}
