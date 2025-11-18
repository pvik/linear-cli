# Linear CLI 

A CLI client to linear.app 

This is primarily aimed at allowing a dev to interact with their issues from a CLI. 
The goal is to improve adoption of Linear.app by devs who have a heavy terminal-focused workflow. 

This allows for devs to quickly create issues for new tasks they may be starting work on, and also automatically create a new local git branch for the linear issue (and switch into it).
This also provides an easy way to switch to git branch for a linear issue.

### Installation 

- Grab pre-built binary from [here](https://github.com/pvik/linear-cli/releases).
- Extract the archive

##### OSX 

- Copy the binary to a location that's included in your `$PATH`
  - If you're using homebrew, you can copy the binary to `/opt/homebrew/bin/`
- OSX requires binaries to be signed, to do this, run the following in your terminal: 

```sh
$ sudo xattr -dr com.apple.quarantine /opt/homebrew/bin/lin
$ sudo codesign -s - --deep --force /opt/homebrew/bin/lin
```

## Usage 

### Configuration

XDG Config location is honored. 
In Linux this should be `~/.config/linear-cli/config.toml` 
In Mac this should be `~/Library/Application Data/linear-cli/config.toml`

This file holds your Linear.app API Key. 
You can create a Personal API Key in Settings > Security & access

```toml
LINEAR_API_KEY = "<key>"
```

#### Project Level Settings 

You can set project level settings in your project directory by creating a `.linear-cli-config.toml`. 

A Sample of this looks like: 

```toml
default_git_head_branch = "main"
default_team = "Engineering"

[default_issue_template]
team = "Engineering"
priority = 3
status = "Todo"
project = "Data"
labels = ["Backend Team"]
```

#### List 

- Issues 
- Teams 
- Projects 

```
NAME:
   lin list - list entities

USAGE:
   lin list [command [command options]]

COMMANDS:
   issue, i, is  list issues
   team, tm      list teams
   project, prj  list projects

OPTIONS:
   --limit int  (default: 50)
   --help, -h   show help
```

##### Examples: 

- To list all you issues under a given team: (all of below are equivalent)
```
lin list issues --team="Engineering" --mine
```

```
lin ls is --tm="Engineering" --mine
```

```
lin l is --tm="Engineering" --mine
```


#### Details 

- Issue 

```
NAME:
   lin detail - show entity detail

USAGE:
   lin detail [command [command options]]

COMMANDS:
   issue, i, is  issue details

OPTIONS:
   --help, -h  show help
```

##### Examples:

- To get details of a specific issue: (all of below are equivalent)

```
lin detail issue --id=ENG-9999
```

```
lin cat issue --id=ENG-9999
```

```
lin cat is --id=ENG-9999
```

```
lin cat i --id=ENG-9999
```



#### Create 

- Issue 

Once an issue is created, linear-cli will try to open the issue in the default browser.

(Note: You can  also set default values for fields like team, label, priority, status in the project level config file)

```
NAME:
   lin new - create a new entity

USAGE:
   lin new [command [command options]]

COMMANDS:
   issue, i, is  issue details

OPTIONS:
   --help, -h  show help
```

##### Examples: 

- Create a new issue ; create a new git-branch for issue and switch to it: (all below are equivalent)

```
lin new issue --team="Engineering" --title="test issue" --label="Backend Team" --label="Improvement" --priority=3 --status="Todo" --project=Data --git-create-branch 
```

```
lin touch is --tm="Engineering" --title="test issue" --label="Backend Team" --label="Improvement" --priority=3 --status="Todo" --prj=Data --g-cb
```

```
lin mk is --tm="Engineering" --title="test issue" --label="Backend Team" --label="Improvement" --priority=3 --status="Todo" --prj=Data --g-cb
```

```
lin mk i --tm="Engineering" --title="test issue" --label="Backend Team" --label="Improvement" --priority=3 --status="Todo" --prj=Data --g-cb
```


### Git 

```
NAME:
   lin git - Linear based helper commands for git repo

USAGE:
   lin git [command [command options]]

COMMANDS:
   switch, sw  checkout to the branch for linear issue (if it already exists)

OPTIONS:
   --help, -h  show help
```

##### Examples 

- Switch branch for linear issue ; Create branch if it doesn't exist: (all below are equivalent)

```
lin git switch --id=ENG-9999 --create-if-not-exists
```

```
lin g sw --id=ENG-9999 --cr
```


## TODO 

- [ ] Cache Teams and Issue Labels locally 
- [ ] Show timestamps in local timezone
