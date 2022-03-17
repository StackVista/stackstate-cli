package cmd

// These variables will be set on build by the ldlflags.
// See the `ldflags` in .goreleaser.yml
var (
	Version string
	Commit  string
	Date    string
	CLIType string
)
