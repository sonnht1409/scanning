package logic

import (
	"context"
	"encoding/base64"

	"github.com/google/go-github/v48/github"
	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/models"
)

func (s ServiceLogic) getRepository(owner, repo string) (*github.Repository, error) {
	resRepo, _, err := s.githubClient.Repositories.Get(context.Background(), owner, repo)
	return resRepo, err
}

func (s ServiceLogic) getAllFilePaths(owner, repo, branch string) ([]*github.TreeEntry, error) {
	tree, _, err := s.githubClient.Git.GetTree(context.Background(), owner, repo, branch, true)
	return tree.Entries, err
}

func (s ServiceLogic) GetRepoContent(owner, repo string) ([]models.GithubFileContent, error) {
	var contents []models.GithubFileContent
	log := common.GetLogger("GetRepoContent", config.Values.Env)
	repository, err := s.getRepository(owner, repo)
	if err != nil {
		log.Errorf("GetRepoContent error %", err.Error())
		return contents, err
	}

	log.Info("GetRepoContent ", *repository.URL, repository.GetDefaultBranch())
	entries, err := s.getAllFilePaths("sonnht1409", "go-crud", repository.GetDefaultBranch())
	if err != nil {
		log.Errorf("GetTree error %", err.Error())
		return contents, err
	}

	for _, entry := range entries {
		log.Infof("Get Entry %s %s\n", *entry.Path, *entry.URL)

		out, err := common.DoHTTPRequest(context.Background(), *entry.URL, "GET", "", "")
		if err != nil {
			log.Infof("Get File Content err %s %s %s\n", *entry.Path, *entry.URL, err.Error())
			continue
		}
		m, _ := out.(map[string]interface{})
		if _, ok := m["content"]; m != nil && ok {
			signBytes, _ := base64.StdEncoding.DecodeString(m["content"].(string))

			strContent := string(signBytes)
			log.Info("file content: ", strContent)
			contents = append(contents, models.GithubFileContent{
				Path:    *entry.Path,
				Url:     *entry.URL,
				Content: strContent,
			})
		}
	}

	return contents, nil
}
