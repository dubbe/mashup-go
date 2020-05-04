package artist

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dubbe/mashup-go/internal/client"
	"github.com/dubbe/mashup-go/pkg/errors"
)

type Musicbrainz struct {
	Client client.HTTPClient
	Url    string
}

func NewMusicbrainz(client client.HTTPClient) *Musicbrainz {
	return &Musicbrainz{
		Client: client,
		Url:    "http://musicbrainz.org/ws/2/artist/%s?&fmt=json&inc=url-rels+release-groups",
	}

}

// Get returns an artist
func (m Musicbrainz) Get(id string) (Artist, error) {
	const op errors.Op = "musicbrainz.Musicbrainz.Get"
	artist := Artist{}

	url := fmt.Sprintf(m.Url, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return artist, errors.E(err, op, "Could not get response from musicbrainz")
	}

	response, err := m.Client.Do(request)
	if err != nil {
		return artist, err
	}
	if response.StatusCode != 200 {
		return artist, errors.E(op, errors.StatusCode(response.StatusCode), fmt.Sprintf("Response from musicbrainz was %d", response.StatusCode))
	}

	json.NewDecoder(response.Body).Decode(&artist)

	return artist, nil
}
