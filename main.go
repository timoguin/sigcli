package main

import "github.com/timoguin/sigcli/cmd"

// compile time build info
var (
	Version   string
	GitCommit string
	BuildTime string
	GoVersion string
)

func main() {
	// Set built-time vars in the cmd package so we can access them in commands
	cmd.Version = Version
	cmd.GitCommit = GitCommit
	cmd.BuildTime = BuildTime
	cmd.GoVersion = GoVersion

	cmd.Execute()
}
