package linear

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Me struct {
	Id    string
	Name  string
	Email string
}

func (l Linear) QueryMe() Me {
	client := l.getClient()

	var queryMe struct {
		Viewer Me
	}

	err := client.Query(context.Background(), &queryMe, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to query teams")
	}

	log.Debug().Any("me", queryMe).Msg("Me")

	return queryMe.Viewer
}
