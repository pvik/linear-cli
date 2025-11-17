package main

import (
	"fmt"
	"pvik/linear-cli/pkg/linear"
	"strings"
)

func detailIssue(issue linear.IssueDetailNode) {
	lbls_lst := []string{}
	for _, lbl := range issue.Labels.Nodes {
		lbls_lst = append(lbls_lst, lbl.Name)
	}
	lbls := strings.Join(lbls_lst, ",")

	fmt.Printf("------------------------\n")
	fmt.Printf("%s: %s\n", issue.Identifier, issue.Title)
	fmt.Printf("%s\n", issue.Url)
	fmt.Printf("------------------------\n")
	fmt.Printf("Priority: %d | Status: %s", issue.Priority, issue.State.Name)
	if issue.Cycle.Name != "" {
		fmt.Printf(" | Cycle: %s", issue.Cycle.Name)
	}
	if issue.Project.Name != "" {
		fmt.Printf(" | Project: %s", issue.Project.Name)
	}
	fmt.Println("")
	fmt.Printf("Labels: %s\n", lbls)
	fmt.Printf("Assigned To: %s | Created By: %s\n", issue.Assignee.Name, issue.Creator.Name)
	fmt.Printf("Updated At: %s | Created At: %s\n", issue.UpdatedAt, issue.CreatedAt)
	fmt.Printf("------------------------\n")
	if issue.Description != "" {
		fmt.Println("Details:")
		fmt.Println(issue.Description)
		fmt.Printf("------------------------\n")
	}
	if len(issue.Comments.Nodes) > 0 {
		fmt.Printf("Comments\n")
		for _, comment := range issue.Comments.Nodes {
			fmt.Printf("------------------------\n")
			fmt.Printf("By %s\n", comment.User.Name)
			fmt.Printf("\tUpdated At: %s | Created At: %s\n", comment.UpdatedAt, comment.CreatedAt)
			fmt.Printf("-----\n")
			fmt.Printf("%s\n", comment.Body)
		}
		fmt.Printf("------------------------\n")
	}
}
