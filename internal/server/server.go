package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/internal/client/coverart"
	"github.com/dubbe/mashup-go/internal/client/description"
	"github.com/dubbe/mashup-go/pkg/errors"
	"github.com/dubbe/mashup-go/pkg/mashup"
	"github.com/google/uuid"
)

type Server struct {
	artistClient       artist.ArtistClient
	descriptionFactory *description.DescriptionFactory
	coverartClient     coverart.CoverartClient
}

func NewServer(artistClient artist.ArtistClient, descriptionFactory *description.DescriptionFactory, coverartClient coverart.CoverartClient) *Server {
	s := &Server{
		artistClient:       artistClient,
		descriptionFactory: descriptionFactory,
		coverartClient:     coverartClient,
	}
	return s
}

func (s *Server) Run() {
	addr := ":8080"
	http.HandleFunc("/", s.handler)
	log.Printf("Server started on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	mbid := strings.TrimPrefix(r.URL.Path, "/")
	_, err := uuid.Parse(mbid)
	if err != nil {
		e := errors.E(err, errors.StatusCode(http.StatusBadRequest), "Not a valid musicbrainz-id")
		handleError(e, w)
		return
	}

	artist, err := s.getArtist(mbid)
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}

func (s *Server) getArtist(mbid string) (mashup.Artist, error) {
	const op errors.Op = "Server.getArtist"
	var desc description.Description
	var albums []mashup.Album

	artist, err := s.artistClient.Get(mbid)
	if err != nil {
		return mashup.Artist{}, errors.E(err, op, "Could not get artist from Musicbrainz")
	}

	relation, err := artist.GetRelation("wikipedia")
	if err != nil {
		relation, err = artist.GetRelation("wikidata")
		if err != nil {
			return mashup.Artist{}, errors.E(err, op, "Could not find discription in neither wikipedia nor wikidata")
		}
	}

	descriptionClient, err := s.descriptionFactory.NewDescriptionClient(relation)
	if err != nil {
		return mashup.Artist{}, errors.E(err, op)
	}

	d := make(chan description.Description)
	de := make(chan error)
	go descriptionClient.Get(d, de)

	a := make(chan []mashup.Album)
	go s.coverartClient.GetMany(artist.Albums, a)

	desc = <-d
	descErr := <-de
	albums = <-a

	if descErr != nil {
		return mashup.Artist{}, errors.E(err, op, "Could not download description")
	}

	return mashup.Artist{
		MBID:        artist.ID,
		Name:        artist.Name,
		Description: string(desc),
		Albums:      albums,
	}, nil
}

func handleError(err error, w http.ResponseWriter) {
	e := err.(*errors.Error)
	sc := e.StatusCodes()
	if len(sc) != 0 {
		w.WriteHeader(int(sc[0]))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(e.Msg))
	e.LogError()
}
