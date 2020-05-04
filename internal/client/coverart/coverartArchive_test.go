package coverart

import (
	"net/http"
	"sort"
	"testing"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/pkg/mashup"
	"github.com/stretchr/testify/assert"
)

func TestMissingCoverart(t *testing.T) {
	coverart := NewCoverartArchive(&http.Client{})

	c := make(chan Coverart)
	go coverart.Get("nope", c)
	art := <-c
	assert.Equal(t, "", art.Image)
}

func TestGetBleachCoverart(t *testing.T) {
	coverart := NewCoverartArchive(&http.Client{})

	c := make(chan Coverart)
	go coverart.Get("f1afec0b-26dd-3db5-9aa1-c91229a74a24", c)
	art := <-c

	assert.Equal(t, "http://coverartarchive.org/release/7d166a44-cfb5-4b08-aacb-6863bbe677d6/1247101964.jpg", art.Image)
}

func TestNirvanaBleachCoverart(t *testing.T) {
	coverart := NewCoverartArchive(&http.Client{})
	albums := []artist.Album{{
		ID:    "f1afec0b-26dd-3db5-9aa1-c91229a74a24",
		Title: "Bleach",
	}, {
		ID:    "1b022e01-4da6-387b-8658-8678046e4cef",
		Title: "Nevermind",
	}}

	c := make(chan []mashup.Album)
	go coverart.GetMany(albums, c)
	art := <-c

	sort.SliceStable(albums, func(i, j int) bool {
		return albums[i].ID < albums[j].ID
	})

	assert.Equal(t, "http://coverartarchive.org/release/7d166a44-cfb5-4b08-aacb-6863bbe677d6/1247101964.jpg", art[0].Image)
	assert.Equal(t, "http://coverartarchive.org/release/a146429a-cedc-3ab0-9e41-1aaf5f6cdc2d/3012495605.jpg", art[1].Image)
}
