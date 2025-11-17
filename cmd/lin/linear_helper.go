package main

import (
	"fmt"
	"pvik/linear-cli/pkg/linear"

	"github.com/rs/zerolog/log"
)

func (a App) getTeamId(teamName string, raiseOnNotFound bool) string {
	c := linear.Linear{ApiKey: a.LinearAPIToken}
	teams := c.QueryTeams(50)

	for _, team := range teams {
		if team.Name == teamName {
			return team.Id
		}
	}

	if raiseOnNotFound {
		log.Fatal().Msg(fmt.Sprintf("Invalid Team (%s) provided.", teamName))
	}

	return ""
}

func (a App) getTeamStateId(teamId string, stateName string, raiseOnNotFound bool) string {
	c := linear.Linear{ApiKey: a.LinearAPIToken}
	states := c.QueryTeamStates(teamId)

	for _, state := range states {
		if state.Name == stateName {
			return state.Id
		}
	}

	if raiseOnNotFound {
		log.Fatal().Msg(fmt.Sprintf("Invalid State (%s) provided for team.", stateName))
	}

	return ""
}

func (a App) getIssueLabelsIds(labels []string, raiseOnNotFound bool) []string {
	c := linear.Linear{ApiKey: a.LinearAPIToken}
	srvLabels := c.QueryIssueLabels(true)

	ret := []string{}

	for _, label := range labels {
		found := false
		for _, srvLabel := range srvLabels {
			if srvLabel.Name == label {
				ret = append(ret, srvLabel.Id)
				found = true
			}
		}

		if !found && raiseOnNotFound {
			log.Fatal().Msg(fmt.Sprintf("Invalid Label (%s) provided.", label))
		}
	}

	return ret
}

func (a App) getMyEmail() string {
	c := linear.Linear{ApiKey: a.LinearAPIToken}
	me := c.QueryMe()

	return me.Email
}

func (a App) getMyId() string {
	c := linear.Linear{ApiKey: a.LinearAPIToken}
	me := c.QueryMe()

	return me.Id
}
