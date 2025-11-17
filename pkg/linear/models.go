package linear

type PageInfo struct {
	HasNextPage bool
	EndCursor   string
}

type TeamNode struct {
	Id   string
	Name string
}

type TeamStateNode struct {
	Id   string
	Name string
}

type IssueLabelNode struct {
	Id   string
	Name string
}

type ProjectNode struct {
	Id          string
	Name        string
	Description string
	Url         string

	CreatedAt string
	UpdatedAt string

	Labels struct {
		Nodes []struct {
			Id   string
			Name string
		}
	}

	Status struct {
		Name string
	}
}

type IssueNode struct {
	Id          string
	Identifier  string
	Title       string
	Description string
	Url         string
	Priority    int

	CreatedAt string
	UpdatedAt string

	Project struct {
		Name string
	}
	State struct {
		Name string
	}
	Assignee struct {
		Id    string
		Name  string
		Email string
	}
}

type IssueDetailNode struct {
	Id          string
	Identifier  string
	Title       string
	Description string
	Url         string
	Priority    int
	BranchName  string

	CreatedAt string
	UpdatedAt string

	Project struct {
		Name string
	}
	State struct {
		Name string
	}
	Creator struct {
		Name string
	}
	Cycle struct {
		Name string
	}
	Assignee struct {
		Id    string
		Name  string
		Email string
	}

	Labels struct {
		Nodes []struct {
			Name string
		}
	}

	Comments struct {
		Nodes []struct {
			Body string

			User struct {
				Name string
			}

			CreatedAt string
			UpdatedAt string
		}
	}
}
