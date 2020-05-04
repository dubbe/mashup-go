package description

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dubbe/mashup-go/internal/client"
	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/pkg/errors"
	"github.com/tidwall/gjson"
)

type Wikipedia struct {
	client     client.HTTPClient
	url        string
	identifier string
	title      string
}

func NewWikipedia(client client.HTTPClient) *Wikipedia {
	return &Wikipedia{
		client:     client,
		url:        "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&exintro=true&redirects=true&titles=%s",
		identifier: "wikipedia",
	}
}

// Get returns an description
func (w *Wikipedia) Get(d chan<- Description, e chan<- error) {
	const op errors.Op = "wikipedia.Wikipedia.Get"

	var description Description

	wikipediaAPIURL := fmt.Sprintf(w.url, url.QueryEscape(w.title))

	request, err := http.NewRequest(http.MethodGet, wikipediaAPIURL, nil)
	if err != nil {
		d <- description
		e <- errors.E(err, op)
		return
	}

	response, err := w.client.Do(request)
	if err != nil {
		d <- description
		e <- errors.E(err, op, errors.StatusCode(response.StatusCode), "Could not download description from Wikipedia")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		d <- description
		e <- errors.E(err, op, "Could not ready body from Wikipedia")
		return
	}

	bodyString := string(body)

	description = Description(gjson.Get(bodyString, "query.pages.*.extract").String())

	d <- description
	e <- nil
	close(d)
	close(e)
	return
}

func (w *Wikipedia) SetRelation(relation artist.Relation) {
	splitString := strings.Split(relation.URL.Resource, "/")
	w.title, _ = url.QueryUnescape(splitString[len(splitString)-1])
}

func (w *Wikipedia) Identifier() string {
	return w.identifier
}
