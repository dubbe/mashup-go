package coverart

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dubbe/mashup-go/internal/client"
	"github.com/dubbe/mashup-go/internal/client/artist"
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

// Get returns an artist
func (c CoverartArchive) Get(id string) (Coverart, error) {
	coverArt := Coverart{}

	url := fmt.Sprintf(c.Url, id)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	if err != nil {
		return coverArt, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return coverArt, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return coverArt, err
	}
	bodyString := string(body)

	coverArt.ID = id
	coverArt.Image = gjson.Get(bodyString, "images.#(front==true).image").String()

	return coverArt, nil
}

// Get returns an artist
func (c CoverartArchive) GetAsync(id string, coverart chan Coverart) error {
	fmt.Println(id)
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
	return nil
}

func (c CoverartArchive) GetMany(albums []artist.Album) ([]Coverart, error) {
	coverArts := []Coverart{}
	for _, album := range albums {
		coverart, _ := c.Get(album.ID)
		coverArts = append(coverArts, Coverart{
			Image: coverart.Image,
			ID:    album.ID,
		})
	}

	return coverArts, nil
}

func (c CoverartArchive) GetManyAsync(albums []artist.Album, coverarts chan []Coverart) {

	coverart := make(chan Coverart)
	for _, album := range albums {
		go c.GetAsync(album.ID, coverart)
	}

	result := make([]Coverart, len(albums))
	for i := range result {
		result[i] = <-coverart
	}

	coverarts <- result
}
