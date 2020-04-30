package server

import (
	"log"
	"net/http"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/internal/client/coverart"
	"github.com/dubbe/mashup-go/internal/client/description"
	"github.com/dubbe/mashup-go/pkg/mashup"
	"github.com/gorilla/mux"
)

type Server struct {
	artistClient       artist.ArtistClient
	descriptionFactory *description.DescriptionFactory
	coverartClient     coverart.CoverartClient
}

func NewServer(artistClient artist.ArtistClient, descriptionFactory *description.DescriptionFactory, coverartClient coverart.CoverartClient) *Server {
	server := &Server{
		artistClient:       artistClient,
		descriptionFactory: descriptionFactory,
		coverartClient:     coverartClient,
	}
	return server
}

func (s *Server) Run() {
	addr := ":8080"
	http.HandleFunc("/{mbid}", s.handler)
	log.Printf("Server started on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mbid := vars["mbid"]
	s.getArtist(mbid)
}

func (s *Server) getArtist(mbid string) mashup.Artist {
	var description description.Description

	artist, err := s.artistClient.Get(mbid)
	if err != nil {
		panic("could not get artist from musicmatch")
	}

	relation, err := artist.GetRelation("wikipedia")
	if err != nil {
		relation, err = artist.GetRelation("wikidata")
		if err != nil {
			panic("could not get description")
		}
	}

	descriptionClient, err := s.descriptionFactory.NewDescriptionClient(relation)
	if err != nil {
		log.Panicf("could not get discription from %s", descriptionClient.Identifier())
	}

	description, _ = descriptionClient.Get()

	_ = description

	albums := []mashup.Album{}
	coverarts, _ := s.coverartClient.GetMany(artist.Albums)
	for _, album := range artist.Albums {
		a := mashup.Album{}
		coverart, err := coverart.GetCoverart(coverarts, album.ID)
		if err != nil {
			a.Image = coverart.Image
		}

		a.ID = album.ID
		a.Title = album.Title
		albums = append(albums, a)
	}

	return mashup.Artist{
		MBID:        artist.ID,
		Name:        artist.Name,
		Description: description.Description,
		Albums:      albums,
	}

}

func (s *Server) getArtistAsync(mbid string) mashup.Artist {
	var description description.Description

	artist, err := s.artistClient.Get(mbid)
	if err != nil {
		panic("could not get artist from musicmatch")
	}

	relation, err := artist.GetRelation("wikipedia")
	if err != nil {
		relation, err = artist.GetRelation("wikidata")
		if err != nil {
			panic("could not get description")
		}
	}

	descriptionClient, err := s.descriptionFactory.NewDescriptionClient(relation)
	if err != nil {
		log.Panicf("could not get discription from %s", descriptionClient.Identifier())
	}

	description, _ = descriptionClient.Get()

	_ = description

	albums := []mashup.Album{}
	c := make(chan []coverart.Coverart)
	go s.coverartClient.GetManyAsync(artist.Albums, c)
	coverarts := <-c

	for _, album := range artist.Albums {
		a := mashup.Album{}
		coverart, err := coverart.GetCoverart(coverarts, album.ID)
		if err != nil {
			a.Image = coverart.Image
		}

		a.ID = album.ID
		a.Title = album.Title
		albums = append(albums, a)
	}

	return mashup.Artist{
		MBID:        artist.ID,
		Name:        artist.Name,
		Description: description.Description,
		Albums:      albums,
	}
}
