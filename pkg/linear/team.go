package linear

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Teams struct {
	Nodes []TeamNode
}

func (l Linear) QueryTeams(limit int) []TeamNode {
	client := l.getClient()

	variables := map[string]interface{}{
		"first": limit,
	}

	var queryTeams struct {
		Teams Teams `graphql:"teams(first: $first)"`
	}

	err := client.Query(context.Background(), &queryTeams, variables)
	if err != nil {
		log.Error().Err(err).Msg("Unable to query teams")
	}

	log.Debug().Any("teams", queryTeams).Msg("Teams")

	return queryTeams.Teams.Nodes
}

func (l Linear) queryTeamIssues(teamIssueVar any, variables map[string]interface{}) {
	client := l.getClient()

	err := client.WithDebug(true).Query(context.Background(), teamIssueVar, variables)
	if err != nil {
		// Handle error.
		log.Fatal().Err(err).Msg("Unable to query Team")
	}
}

func (l Linear) QueryTeamIssues(teamId string) []IssueNode {
	variables := map[string]interface{}{
		"id": teamId,
	}

	var queryTeamIssuess struct {
		Team struct {
			Id   string
			Name string

			Issues struct {
				Nodes []IssueNode
			} `graphql:"issues(orderBy: updatedAt)"`
		} `graphql:"team(id: $id)"`
	}

	l.queryTeamIssues(&queryTeamIssuess, variables)

	// log.Debug().Any("team-issues", queryTeamIssuess).Msg("Team Issues")

	return queryTeamIssuess.Team.Issues.Nodes
}

func (l Linear) QueryTeamIssuesOpen(teamId string, limit int) []IssueNode {
	variables := map[string]interface{}{
		"id":    teamId,
		"first": limit,
	}

	var queryTeamIssuess struct {
		Team struct {
			Id   string
			Name string

			Issues struct {
				Nodes []IssueNode
			} `graphql:"issues(orderBy: updatedAt, first: $first, filter: {and: [{state: {name: {neq: \"Done\"}}}, {state: {name: {neq: \"Canceled\"}}}, {state: {name: {neq: \"Duplicate\"}}}]})"`
		} `graphql:"team(id: $id)"`
	}

	l.queryTeamIssues(&queryTeamIssuess, variables)

	// log.Debug().Any("team-issues", queryTeamIssuess).Msg("Team Issues")

	return queryTeamIssuess.Team.Issues.Nodes
}

func (l Linear) QueryTeamIssuesByAssigned(teamId string, assignedEmail string) []IssueNode {
	variables := map[string]interface{}{
		"id":            teamId,
		"assignedEmail": assignedEmail,
	}

	var queryTeamIssuess struct {
		Team struct {
			Id   string
			Name string

			Issues struct {
				Nodes []IssueNode
			} `graphql:"issues(orderBy: updatedAt, filter: { assignee: { email: { eq: $assignedEmail } } })"`
		} `graphql:"team(id: $id)"`
	}

	l.queryTeamIssues(&queryTeamIssuess, variables)

	// log.Debug().Any("team-issues", queryTeamIssuess).Msg("Team Issues")

	return queryTeamIssuess.Team.Issues.Nodes
}

func (l Linear) QueryTeamIssuesByAssignedOpen(teamId string, assignedEmail string, limit int) []IssueNode {
	variables := map[string]interface{}{
		"id":            teamId,
		"assignedEmail": assignedEmail,
		"first":         limit,
	}

	var queryTeamIssuess struct {
		Team struct {
			Id   string
			Name string

			Issues struct {
				Nodes []IssueNode
			} `graphql:"issues(orderBy: updatedAt, first: $first, filter: {and: [{assignee: {email: {eq: $assignedEmail}}}, {state: {name: {neq: \"Done\"}}}, {state: {name: {neq: \"Canceled\"}}}, {state: {name: {neq: \"Duplicate\"}}}]})"`
		} `graphql:"team(id: $id)"`
	}

	l.queryTeamIssues(&queryTeamIssuess, variables)

	// log.Debug().Any("team-issues", queryTeamIssuess).Msg("Team Issues")

	return queryTeamIssuess.Team.Issues.Nodes
}

func (l Linear) QueryTeamStates(teamId string) []TeamStateNode {
	variables := map[string]interface{}{
		"id": teamId,
	}

	var queryTeamStates struct {
		Team struct {
			Id   string
			Name string

			States struct {
				Nodes []TeamStateNode
			} `graphql:"states(orderBy: updatedAt)"`
		} `graphql:"team(id: $id)"`
	}

	l.queryTeamIssues(&queryTeamStates, variables)

	// log.Debug().Any("team-issues", queryTeamIssuess).Msg("Team Issues")

	return queryTeamStates.Team.States.Nodes
}
