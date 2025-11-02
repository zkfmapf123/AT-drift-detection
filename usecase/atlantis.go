package usecase

type AtlantisRequest struct {
	GithubToken   string
	GithubRepoRef string

	AtlantisURL        string
	AtlantisToken      string
	AtlantisRepository string
	AtlantisConfigFile string
}

type APIHealthResponse struct {
	Status string `json:"status"`
}
