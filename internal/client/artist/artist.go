package artist

import "errors"

// Album is an album
type Artist struct {
	ID        string
	Name      string
	Relations []Relation
	Albums    []Album `json:"release-groups"`
}

type Relation struct {
	Type string
	URL  URL
}

type Album struct {
	Title string
	ID    string
}

type URL struct {
	Resource string
}

type ArtistClient interface {
	Get(id string) (Artist, error)
}

func (a Artist) FilterRelations(s string) []Relation {
	rsf := make([]Relation, 0)
	for _, r := range a.Relations {
		if r.Type == s {
			rsf = append(rsf, r)
		}
	}
	return rsf
}

func (a Artist) GetRelation(s string) (Relation, error) {
	relations := a.FilterRelations(s)
	if len(relations) == 1 {
		return relations[0], nil
	}

	return Relation{}, errors.New("Could not get just one relation")
}
