package linear

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Issue struct {
	Id          string
	Title       string
	Description string
}

type IssueMin struct {
	Id    string
	Title string
	// Description string
}

func (l Linear) QueryIssue(issueId string) IssueDetailNode {
	client := l.getClient()

	variables := map[string]interface{}{
		"id": issueId,
	}

	var queryIssue struct {
		Issue IssueDetailNode `graphql:"issue(id: $id)"`
	}

	err := client.Query(context.Background(), &queryIssue, variables)
	if err != nil {
		// Handle error.
		log.Fatal().Err(err).Msg("Unable to query issue")
	}

	return queryIssue.Issue
}

func (l Linear) CreateIssue(
	teamId string,
	title string,
	description string,
	assigneeId string,
	stateId string,
	priority int,
	labelsIds []string,
) IssueDetailNode {
	client := l.getClient()

	variables := map[string]interface{}{
		"teamId":      teamId,
		"title":       title,
		"description": description,
		"assigneeId":  assigneeId,
		"stateId":     stateId,
		"priority":    priority,
		"labelIds":    labelsIds,
	}

	var mutateIssue struct {
		IssueCreate struct {
			Success bool
			Issue   IssueDetailNode
		} `graphql:"issueCreate(input: {title: $title, description: $description, teamId: $teamId, assigneeId: $assigneeId, stateId: $stateId, priority: $priority, labelIds: $labelIds})"`
	}

	err := client.WithDebug(true).Mutate(context.Background(), &mutateIssue, variables)
	if err != nil {
		// Handle error.
		log.Fatal().Err(err).Msg("Unable to create issue")
	}

	return mutateIssue.IssueCreate.Issue
}

func (l Linear) queryIssueLabelsAfter(after string) ([]IssueLabelNode, PageInfo) {
	client := l.getClient()

	variables := map[string]interface{}{
		"after": after,
	}

	var queryIssueLabels struct {
		IssueLabels struct {
			Nodes    []IssueLabelNode
			PageInfo PageInfo
		} `graphql:"issueLabels(orderBy: updatedAt, first: 50, after: $after)"`
	}

	err := client.Query(context.Background(), &queryIssueLabels, variables)
	if err != nil {
		// Handle error.
		log.Fatal().Err(err).Msg("Unable to query issue labels")
	}

	return queryIssueLabels.IssueLabels.Nodes, queryIssueLabels.IssueLabels.PageInfo
}

func (l Linear) QueryIssueLabels(all bool) []IssueLabelNode {
	client := l.getClient()

	var queryIssueLabels struct {
		IssueLabels struct {
			Nodes    []IssueLabelNode
			PageInfo PageInfo
		} `graphql:"issueLabels(orderBy: updatedAt, first: 50)"`
	}

	err := client.WithDebug(true).Query(context.Background(), &queryIssueLabels, nil)
	if err != nil {
		// Handle error.
		log.Fatal().Err(err).Msg("Unable to query issue labels")
	}

	hasNext := queryIssueLabels.IssueLabels.PageInfo.HasNextPage
	cursor := queryIssueLabels.IssueLabels.PageInfo.EndCursor

	if all {
		for hasNext {
			issueLabelNodes, pgInfo := l.queryIssueLabelsAfter(cursor)

			hasNext = pgInfo.HasNextPage
			cursor = pgInfo.EndCursor

			queryIssueLabels.IssueLabels.Nodes = append(queryIssueLabels.IssueLabels.Nodes, issueLabelNodes...)
		}
	}

	return queryIssueLabels.IssueLabels.Nodes
}
