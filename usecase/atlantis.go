package usecase

type AtlantisRequest struct {
	GithubToken string

	AtlantisURL        string
	AtlantisToken      string
	AtlantisRepository string
	AtlantisConfigFile string
}
