package payloads

import "github.com/defendops/bedro-confuser/pkg/utils/requester"

type PayloadType string

type GithubAPI struct {
	BaseURL 	string
	Repository GitHubRepository
	Requester requester.HTTPRequester
}

type GitHubRepository struct {
	Owner      string
	Repo       string
	Branch     string
}

type GitHubFile struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	SHA         string `json:"sha"`
	URL         string `json:"url"`
	GitURL      string `json:"git_url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
	Links       Links  `json:"_links"`
}

type Links struct {
	Git  string `json:"git"`
	Self string `json:"self"`
	HTML string `json:"html"`
}