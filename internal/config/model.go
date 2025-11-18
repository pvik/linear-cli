package config

type AppConfig struct {
	LinearAPIToken string `toml:"LINEAR_API_KEY"`
}

type ProjectConfig struct {
	DefaultGitHeadBranch string `toml:"default_git_head_branch"`
	DefaultTeam          string `toml:"default_team"`
	DefaultIssueTemplate struct {
		Team     string   `toml:"team"`
		Priority int      `toml:"priority"`
		Status   string   `toml:"status"`
		Project  string   `toml:"project"`
		Labels   []string `toml:"labels"`
	} `toml:"default_issue_template"`
}
