package linear

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (l Linear) QueryProjects() []ProjectNode {
	client := l.getClient()

	var queryProjects struct {
		Projects struct {
			Nodes []ProjectNode
		} `graphql:"projects(orderBy: updatedAt, filter: {and: [{status: {name: {neq: \"Completed\"}}}]})"`
	}

	err := client.Query(context.Background(), &queryProjects, nil)
	if err != nil {
		log.Error().Err(err).Msg("Unable to query teams")
	}

	log.Debug().Any("projects", queryProjects).Msg("Projects")

	return queryProjects.Projects.Nodes
}
