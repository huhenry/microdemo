package version

var (
	// commitFromGit is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	commitSHA string
	// versionFromGit is a constant representing the version tag that
	// generated this build. It should be set during build via -ldflags.
	latestVersion string
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	date string
	// go version build with. It should be updated when go mod version changed.
	golangVersion = "1.14"
)

// Info holds build information
type Info struct {
	GitCommit     string `json:"gitCommit"`
	GitVersion    string `json:"gitVersion"`
	BuildDate     string `json:"buildDate"`
	GolangVersion string `json:"golangVersion"`
}

// Get creates and initialized Info object
func Get() Info {
	return Info{
		GitCommit:     commitSHA,
		GitVersion:    latestVersion,
		BuildDate:     date,
		GolangVersion: golangVersion,
	}
}
