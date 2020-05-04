package coverart

import (
	"errors"

	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/pkg/mashup"
)

type Coverart struct {
	ID    string
	Image string
}

type CoverartClient interface {
	Get(string, chan<- Coverart)
	GetMany([]artist.Album, chan<- []mashup.Album)
}

func FilterCoverart(coverarts []Coverart, s string) []Coverart {
	rsf := make([]Coverart, 0)
	for _, c := range coverarts {
		if c.ID == s {
			rsf = append(rsf, c)
		}
	}
	return rsf
}

func GetCoverart(c []Coverart, s string) (Coverart, error) {
	relations := FilterCoverart(c, s)
	if len(relations) == 1 {
		return relations[0], nil
	}

	return Coverart{}, errors.New("Could not get just one relation")
}
