package main

import (
	"context"
	"os"
	"pvik/linear-cli/pkg/git"
	"pvik/linear-cli/pkg/linear"

	"github.com/jasonlovesdoggo/gopen"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

func (a App) ParseCLIParams() {
	log.Debug().Msg("Parsing CLI Params")

	cmd := &cli.Command{
		UseShortOptionHandling: true,

		Flags: []cli.Flag{
			// &cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
		},

		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list entities",
				Aliases: []string{"l", "ls"},

				Flags: []cli.Flag{
					&cli.IntFlag{Name: "limit", Value: 50},
				},

				Commands: []*cli.Command{
					{
						Name:    "issue",
						Usage:   "list issues",
						Aliases: []string{"i", "is"},
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "mine", Usage: "Only list issues assigned to me"},
							&cli.StringFlag{Name: "team", Aliases: []string{"tm"}, Usage: "Issues for team"},
						},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							log.Debug().Msg("listing issues")
							log.Debug().Any("team:", cmd.String("team")).Msg("Team cli arg")

							teamName := cmd.String("team")
							if teamName == "" {
								if a.ProjectConfig.DefaultTeam == "" {
									log.Fatal().Msg("Please provide valid Team")
								}
								teamName = a.ProjectConfig.DefaultTeam
							}
							mine := cmd.Bool("mine")
							limit := cmd.Int("limit")

							teamId := a.getTeamId(teamName, true)

							c := linear.Linear{ApiKey: a.LinearAPIToken}

							var issues []linear.IssueNode
							if mine {
								my_email := a.getMyEmail()
								issues = c.QueryTeamIssuesByAssignedOpen(teamId, my_email, limit)
							} else {
								issues = c.QueryTeamIssuesOpen(teamId, limit)
							}

							tabulateIssues(issues)

							return nil
						},
					},

					{
						Name:    "team",
						Usage:   "list teams",
						Aliases: []string{"tm"},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							log.Debug().Msg("listing teams")

							limit := cmd.Int("limit")

							c := linear.Linear{ApiKey: a.LinearAPIToken}
							teams := c.QueryTeams(limit)

							tabulateTeams(teams)

							return nil
						},
					},

					{
						Name:    "project",
						Usage:   "list projects",
						Aliases: []string{"prj"},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							log.Debug().Msg("listing projects")

							c := linear.Linear{ApiKey: a.LinearAPIToken}
							projects := c.QueryProjects()

							tabulateProjects(projects)

							return nil
						},
					},
				},
			},

			{
				Name:    "detail",
				Usage:   "show entity detail",
				Aliases: []string{"cat"},

				Flags: []cli.Flag{},

				Commands: []*cli.Command{
					{
						Name:    "issue",
						Usage:   "issue details",
						Aliases: []string{"i", "is"},
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id", Required: true},
						},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							log.Debug().Msg("detail issue")

							issueId := cmd.String("id")

							c := linear.Linear{ApiKey: a.LinearAPIToken}

							issue := c.QueryIssue(issueId)

							detailIssue(issue)
							return nil
						},
					},
				},
			},

			{
				Name:    "new",
				Usage:   "create a new entity",
				Aliases: []string{"touch", "mk"},

				Flags: []cli.Flag{},

				Commands: []*cli.Command{
					{
						Name:    "issue",
						Usage:   "issue details",
						Aliases: []string{"i", "is"},
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "title", Required: true},
							&cli.StringFlag{Name: "team", Aliases: []string{"tm"}},
							&cli.StringSliceFlag{Name: "label", Value: a.ProjectConfig.DefaultIssueTemplate.Labels},
							&cli.IntFlag{Name: "priority", Value: a.ProjectConfig.DefaultIssueTemplate.Priority},
							&cli.StringFlag{Name: "status", Value: a.ProjectConfig.DefaultIssueTemplate.Status},
							&cli.BoolFlag{Name: "git-create-branch", Aliases: []string{"g-cb"}, Usage: "Create a new git branch for Issue, and switch to it."},
						},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							log.Debug().Msg("create new issue")

							title := cmd.String("title")

							teamName := cmd.String("team")
							if teamName == "" {
								if a.ProjectConfig.DefaultTeam == "" {
									log.Fatal().Msg("Please provide valid Team")
								}
								teamName = a.ProjectConfig.DefaultTeam
							}
							teamId := a.getTeamId(teamName, true)

							priority := cmd.Int("priority")

							status := cmd.String("status")
							if cmd.IsSet("priority") {
								priority = cmd.Int("priority")
							}
							stateId := ""
							if status != "" {
								stateId = a.getTeamStateId(teamId, status, true)
							}

							labels := cmd.StringSlice("label")
							labelsIds := []string{}
							if len(labels) > 0 {
								labelsIds = append(labelsIds, a.getIssueLabelsIds(labels, true)...)
							}

							myId := a.getMyId()

							c := linear.Linear{ApiKey: a.LinearAPIToken}
							issue := c.CreateIssue(teamId, title, "", myId, stateId, priority, labelsIds)

							log.Info().Msgf("Created new issue: %s", issue.Identifier)
							detailIssue(issue)

							// git branch create
							createBranch := cmd.Bool("git-create-branch")

							if createBranch {
								sourceBranch := a.ProjectConfig.DefaultGitHeadBranch
								if sourceBranch == "" {
									sourceBranch = "master"
								}
								git.CreateBranch(issue.BranchName, sourceBranch, true)
							}

							err := gopen.Open(issue.Url)
							if err != nil {
								log.Debug().Msg("Unable to open issue URL")
							}

							return nil
						},
					},
				},
			},

			{
				Name:    "git",
				Usage:   "Linear based helper commands for git repo",
				Aliases: []string{"g"},

				Flags: []cli.Flag{},

				Commands: []*cli.Command{
					{
						Name:    "switch",
						Usage:   "checkout to the branch for linear issue (if it already exists)",
						Aliases: []string{"sw"},
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "id", Aliases: []string{"issue-id"}, Required: true},
							&cli.BoolFlag{Name: "create-if-not-exists", Aliases: []string{"cr"}},
						},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							issueId := cmd.String("id")

							c := linear.Linear{ApiKey: a.LinearAPIToken}

							issue := c.QueryIssue(issueId)

							branchExists := git.CheckBranchExists(issue.BranchName)
							if branchExists {
								git.SwitchBranch(issue.BranchName)
							} else {
								if cmd.IsSet("create-if-not-exists") && cmd.Bool("create-if-not-exists") {
									sourceBranch := a.ProjectConfig.DefaultGitHeadBranch
									if sourceBranch == "" {
										sourceBranch = "master"
									}
									git.CreateBranch(issue.BranchName, sourceBranch, true)
								} else {
									log.Fatal().Msgf("Branch (%s) does not exist. Pass in --create-if-not-exists flag to create branch", issue.BranchName)
								}
							}
							return nil
						},
					},

					{
						Name:    "zbranch",
						Usage:   "(for testing only) create new branch and checkout",
						Aliases: []string{"b"},
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name"},
						},

						Action: func(ctx context.Context, cmd *cli.Command) error {
							branchName := cmd.String("name")
							git.CreateBranch(branchName, "main", true)
							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("Unable to parse CLI args")
	}
}
