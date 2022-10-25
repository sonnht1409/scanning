package logic

import (
	"context"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/models"
)

func (s ServiceLogic) getRepository(owner, repo string) (*github.Repository, error) {
	resRepo, _, err := s.githubClient.Repositories.Get(context.Background(), owner, repo)
	return resRepo, err
}

/*
Get all git tree entries
Ex:
"tree": [
{
"path": "README.md",
"mode": "100644",
"type": "blob",
"sha": "759cb15cd5b7d30186a30d61ba56490dbfca5456",
"size": 733,
"url": "https://api.github.com/repos/sonnht1409/go-crud/git/blobs/759cb15cd5b7d30186a30d61ba56490dbfca5456"
}
]
*/
func (s ServiceLogic) getAllFilePaths(owner, repo, branch string) ([]*github.TreeEntry, error) {
	tree, _, err := s.githubClient.Git.GetTree(context.Background(), owner, repo, branch, true)
	if tree != nil {
		return tree.Entries, err
	}
	return nil, err
}

/*
Get file content by it sha commit
*/
func (s ServiceLogic) getFileContent(owner, repo, sha string) (*github.Blob, error) {
	blob, _, err := s.githubClient.Git.GetBlob(context.Background(), owner, repo, sha)
	return blob, err
}

func (s ServiceLogic) GetRepoContent(owner, repo string) ([]models.GithubFileContent, error) {
	var contents []models.GithubFileContent
	log := common.GetLogger("GetRepoContent", config.Values.Env)
	repository, err := s.getRepository(owner, repo)
	if err != nil {
		log.Errorf("GetRepoContent error %s %s", owner, repo)
		return contents, err
	}

	if repository == nil {
		log.Warnf("Repository is nil %s %s", owner, repo)
		return contents, err
	}

	log.Info("GetRepoContent ", *repository.URL, repository.GetDefaultBranch())
	entries, err := s.getAllFilePaths("sonnht1409", "go-crud", repository.GetDefaultBranch())
	if err != nil {
		log.Errorf("GetGitTree error %s %s %s", owner, repo, repository.GetDefaultBranch())
		return contents, err
	}

	for _, entry := range entries {
		if *entry.Type != "blob" {
			continue
		}
		log.Infof("Get Entry %s %s\n", *entry.Path, *entry.URL)
		blob, err := s.getFileContent(owner, repo, *entry.SHA)
		if err != nil {
			log.Errorf("Get File Content err %s %s \n", *entry.Path, *entry.URL)
			continue
		}
		if blob != nil && blob.Content != nil {
			contents = append(contents, models.GithubFileContent{
				Path:    *entry.Path,
				Url:     *entry.URL,
				Content: *blob.Content,
			})
		}
	}

	return contents, nil
}

/*
Assume that public repos have format like: https://github.com/guardrailsio/backend-engineer-challenge
First, clear https://
Then split github.com/guardrailsio/backend-engineer-challenge into [github.com,guardrailsio,backend-engineer-challenge]
Then return guardrailsio
*/
func (s ServiceLogic) GetRepoOwner(repoUrl string) string {
	url := strings.ReplaceAll(repoUrl, "https://", "")
	url = strings.ReplaceAll(url, "http://", "")
	strs := strings.Split(url, "/")
	if len(strs) < 2 {
		return ""
	}
	return strs[1]
}
