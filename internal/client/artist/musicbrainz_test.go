package artist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNirvana(t *testing.T) {
	musicbrainz := NewMusicbrainz(&http.Client{})

	artist, err := musicbrainz.Get("5b11f4ce-a62d-471e-81fc-a69a8278c7da")
	assert.Equal(t, nil, err, "Error should be nil")
	assert.Equal(t, "Nirvana", artist.Name, "name should be Nirvana")
	assert.Equal(t, "https://www.wikidata.org/wiki/Q11649", artist.FilterRelations("wikidata")[0].URL.Resource)
	assert.Equal(t, 25, len(artist.Albums))
}

func TestFilterRelations(t *testing.T) {

	file, _ := ioutil.ReadFile("../../../test/data/nirvana.json")

	artist := Artist{}
	err := json.Unmarshal([]byte(file), &artist)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://www.wikidata.org/wiki/Q11649", artist.FilterRelations("wikidata")[0].URL.Resource)
}
