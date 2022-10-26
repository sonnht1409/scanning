package logic

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/models"
	"golang.org/x/oauth2"
)

var _ IServiceLogic = (*ServiceLogic)(nil)

//go:generate mockery --name IServiceLogic --inpackage
type IServiceLogic interface {
	GetRepoContent(string, string) ([]models.GithubFileContent, error)
	GetRepoOwner(repoUrl string) string
	CheckRule(content string, rule models.RegexRule) (bool, []int)
	getRepository(owner, repo string) (*github.Repository, error)
	getAllFilePaths(owner, repo, branch string) ([]*github.TreeEntry, error)
	getFileContent(owner, repo, sha string) (*github.Blob, error)
}

type ServiceLogic struct {
	githubClient *github.Client
}

func NewServiceLogic() ServiceLogic {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Values.AccessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	return ServiceLogic{
		githubClient: client,
	}
}
