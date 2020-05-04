package artist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dubbe/mashup-go/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetNirvana(t *testing.T) {
	musicbrainz := NewMusicbrainz(&http.Client{})

	artist, err := musicbrainz.Get("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Nirvana", artist.Name)
	assert.Equal(t, "https://www.wikidata.org/wiki/Q11649", artist.FilterRelations("wikidata")[0].URL.Resource)
	assert.Equal(t, 25, len(artist.Albums))
}

func TestGetEbbaGron(t *testing.T) {
	musicbrainz := NewMusicbrainz(&http.Client{})

	artist, err := musicbrainz.Get("4a41bce5-f225-4e0e-91d8-bf49ebbd83c2")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Ebba Grön", artist.Name)
	assert.Equal(t, "https://en.wikipedia.org/wiki/Ebba_Gr%C3%B6n", artist.FilterRelations("wikipedia")[0].URL.Resource)
	assert.Equal(t, 13, len(artist.Albums))
}

func TestGetMissing(t *testing.T) {
	musicbrainz := NewMusicbrainz(&http.Client{})

	_, err := musicbrainz.Get("000")
	assert.Equal(t, errors.Op("musicbrainz.Musicbrainz.Get"), err.(*errors.Error).Op, "Error should be nil")
	assert.Equal(t, errors.StatusCode(400), err.(*errors.Error).StatusCode, "Error should be nil")
}

func TestFilterRelations(t *testing.T) {

	file, _ := ioutil.ReadFile("../../../test/data/nirvana.json")

	artist := Artist{}
	err := json.Unmarshal([]byte(file), &artist)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://www.wikidata.org/wiki/Q11649", artist.FilterRelations("wikidata")[0].URL.Resource)
}
