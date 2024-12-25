package payloads

import (
	"encoding/json"
	"fmt"
	"io"

	utilsRequester "github.com/defendops/bedro-confuser/pkg/utils/requester"
)

func NewGithubAPI() *GithubAPI {
	requester := utilsRequester.HTTPRequester{}

	return &GithubAPI{
		BaseURL: "https://api.github.com/",
		Repository: GitHubRepository{
			Owner: "defendops",
			Repo: "bedro-confusion",
			Branch: "payloads",
		},
		Requester: requester,
	}
}

func (gh *GithubAPI) GetFile(filePath string) (GitHubFile, error){
	request := utilsRequester.HTTPRequest{
		BaseURL: gh.BaseURL,
		Endpoint: fmt.Sprintf("repos/%s/%s/contents/%s?ref=%s", gh.Repository.Owner, gh.Repository.Repo, filePath, gh.Repository.Branch),
		Method: "GET",
		IsJson: false,
		Body: "",
		Headers: GithubHeaders,
	}

	resp, err := gh.Requester.PerformRequest(request)
	if err != nil {
		return GitHubFile{}, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return GitHubFile{}, err
	}
	
	var file GitHubFile
	err = json.Unmarshal(response, &file)
	if err != nil {
		return GitHubFile{}, err
	}

	return file, nil
}

func (gh *GithubAPI) ListFolder(directoryPath string, files_only bool) ([]GitHubFile, error){
	var defaultEndpoint string
	if directoryPath == "" {
		defaultEndpoint = fmt.Sprintf("repos/%s/%s/contents?ref=%s", gh.Repository.Owner, gh.Repository.Repo, gh.Repository.Branch)
	}else{
		defaultEndpoint = fmt.Sprintf("repos/%s/%s/contents/%s?ref=%s", gh.Repository.Owner, gh.Repository.Repo, directoryPath, gh.Repository.Branch)
	}

	request := utilsRequester.HTTPRequest{
		BaseURL: gh.BaseURL,
		Endpoint: defaultEndpoint,
		Method: "GET",
		IsJson: false,
		Body: "",
		Headers: GithubHeaders,
	}

	resp, err := gh.Requester.PerformRequest(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var formattedResponse []GitHubFile
	err = json.Unmarshal(response, &formattedResponse)
	if err != nil {
		return nil, err
	}

	if files_only{
		var FilesResponse []GitHubFile
		for _, file := range formattedResponse{
			if file.Type == "dir"{
				files, err := gh.ListFolder(file.Path, files_only)
				if err != nil{
					continue
				}
				
				FilesResponse = append(FilesResponse, files...)
			}else if file.Type == "file"{
				fileContent, err := gh.GetFile(file.Path)
				if err != nil{
					fileContent = file
				}

				FilesResponse = append(FilesResponse, fileContent)
			}
		}
		return FilesResponse, nil
	}

	return formattedResponse, nil
}

func (gh *GithubAPI) GetPayloadFiles(payload PayloadType) ([]GitHubFile) {
	files, err := gh.ListFolder(string(payload), true)
	if err != nil{
		return nil
	}
	
	return files
}