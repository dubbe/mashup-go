package coverart

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dubbe/mashup-go/internal/client"
	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/pkg/mashup"
	"github.com/tidwall/gjson"
)

type CoverartArchive struct {
	Client client.HTTPClient
	Url    string
}

func NewCoverartArchive(client client.HTTPClient) *CoverartArchive {
	return &CoverartArchive{
		Client: client,
		Url:    "http://coverartarchive.org/release-group/%s",
	}

}

func (c CoverartArchive) Get(id string, coverart chan<- Coverart) {
	coverArt := Coverart{}

	url := fmt.Sprintf(c.Url, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	if err != nil {
		panic("error")
	}

	response, err := c.Client.Do(request)
	if err != nil {
		panic("error")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("error")
	}
	bodyString := string(body)

	coverArt.ID = id
	coverArt.Image = gjson.Get(bodyString, "images.#(front==true).image").String()

	coverart <- coverArt
	return
}

func (ca CoverartArchive) GetMany(albums []artist.Album, albumsChannel chan<- []mashup.Album) {

	coverart := make(chan Coverart)
	for _, album := range albums {
		go ca.Get(album.ID, coverart)
	}

	results := make([]Coverart, len(albums))
	for i := range results {
		results[i] = <-coverart
	}

	returnAlbums := []mashup.Album{}
	for _, album := range albums {
		a := mashup.Album{
			ID:    album.ID,
			Title: album.Title,
		}
		c, err := GetCoverart(results, album.ID)
		if err == nil {
			a.Image = c.Image
		}

		returnAlbums = append(returnAlbums, a)
	}

	albumsChannel <- returnAlbums
	return
}
