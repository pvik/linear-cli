package main

import (
	"fmt"
	"os"
	"pvik/linear-cli/pkg/linear"
	"strings"
	"text/tabwriter"
)

func tabulateTeams(teams []linear.TeamNode) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintf(w, "ID\tName\n")
	fmt.Fprintf(w, "--\t----\n")
	for _, team := range teams {
		fmt.Fprintf(w, "%s\t%s\n", team.Id, team.Name)
	}

	w.Flush()
}

func tabulateIssues(issues []linear.IssueNode) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprint(w, "ID\tUpdated At\tAssigned\tStatus\tProject\t \n")
	fmt.Fprint(w, "--\t----------\t--------\t------\t-------\t \n")
	for _, issue := range issues {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", issue.Identifier, issue.UpdatedAt, issue.Assignee.Name, issue.State.Name, issue.Project.Name, issue.Title)
	}

	w.Flush()
}

func tabulateProjects(issues []linear.ProjectNode) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintf(w, "Name\tUpdated At\tStatus\tLabels\n")
	fmt.Fprintf(w, "----\t----------\t------\t------\n")
	for _, issue := range issues {
		lbls_lst := []string{}
		for _, lbl := range issue.Labels.Nodes {
			lbls_lst = append(lbls_lst, lbl.Name)
		}
		lbls := strings.Join(lbls_lst, ",")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", issue.Name, issue.UpdatedAt, issue.Status.Name, lbls)
	}

	w.Flush()
}
