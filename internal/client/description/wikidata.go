package description

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dubbe/mashup-go/internal/client"
	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/tidwall/gjson"
)

type Wikidata struct {
	client     client.HTTPClient
	url        string
	identifier string
	wikipedia  *Wikipedia
	id         string
}

func NewWikidata(client client.HTTPClient, wikipedia *Wikipedia) *Wikidata {
	return &Wikidata{
		client:     client,
		url:        "https://www.wikidata.org/w/api.php?action=wbgetentities&ids=%s&format=json&props=sitelinks",
		identifier: "wikidata",
		wikipedia:  wikipedia,
	}

}

// Get returns an description
func (w *Wikidata) Get() (Description, error) {
	description := Description{}

	wikidataAPIURL := fmt.Sprintf(w.url, url.QueryEscape(w.id))

	request, err := http.NewRequest(http.MethodGet, wikidataAPIURL, nil)
	if err != nil {
		return description, err
	}

	response, err := w.client.Do(request)
	if err != nil {
		return description, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return description, err
	}
	bodyString := string(body)

	// pages := gjson.Get(bodyString, "query.pages")
	title := gjson.Get(bodyString, "entities.*.sitelinks.enwiki.title").String()

	w.wikipedia.title = title
	description, _ = w.wikipedia.Get()

	return description, nil
}

func (w *Wikidata) SetRelation(relation artist.Relation) {
	splitString := strings.Split(relation.URL.Resource, "/")
	w.id = splitString[len(splitString)-1]
}

func (w *Wikidata) Identifier() string {
	return w.identifier
}
