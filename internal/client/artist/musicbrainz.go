package artist

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dubbe/mashup-go/internal/client"
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
	artist := Artist{}

	url := fmt.Sprintf(m.Url, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return artist, err
	}

	response, err := m.Client.Do(request)
	if err != nil {
		return artist, err
	}

	json.NewDecoder(response.Body).Decode(&artist)

	return artist, nil
}
