package description

import (
	"github.com/dubbe/mashup-go/internal/client/artist"
	"github.com/dubbe/mashup-go/pkg/errors"
)

type Description string

// Interface for possible descriptiongetter
type DescriptionClient interface {
	Identifier() string
	Get(chan<- Description, chan<- error)
	SetRelation(artist.Relation)
}

type DescriptionFactory struct {
	clients []DescriptionClient
}

func CreateDescriptionFactory(clients ...DescriptionClient) *DescriptionFactory {
	return &DescriptionFactory{
		clients: clients,
	}
}

func (d DescriptionFactory) NewDescriptionClient(relation artist.Relation) (DescriptionClient, error) {
	descriptionClients := d.FilterDescriptionClients(relation.Type)
	if len(descriptionClients) == 1 {
		descriptionClient := descriptionClients[0]
		descriptionClient.SetRelation(relation)
		return descriptionClient, nil
	}

	return nil, errors.E("Could not get one descriptionClient only")
}

func (d DescriptionFactory) FilterDescriptionClients(s string) []DescriptionClient {
	dc := make([]DescriptionClient, 0)
	for _, c := range d.clients {
		if c.Identifier() == s {
			dc = append(dc, c)
		}
	}
	return dc
}
