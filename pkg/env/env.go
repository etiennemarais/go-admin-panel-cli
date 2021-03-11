package env

// Env defines the environment variables that are parsed during startup
type Env struct {
	GithubToken string `envconfig:"GITHUB_TOKEN" required:"true"`
}
