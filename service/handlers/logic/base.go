package logic

import (
	"net/http"

	"github.com/google/go-github/v48/github"
	"github.com/sonnht1409/scanning/service/models"
)

var _ IServiceLogic = (*ServiceLogic)(nil)

//go:generate mockery --name IServiceLogic --inpackage
type IServiceLogic interface {
	GetRepoContent(string, string) ([]models.GithubFileContent, error)
}

type ServiceLogic struct {
	githubClient *github.Client
}

func NewServiceLogic() ServiceLogic {
	return ServiceLogic{
		githubClient: github.NewClient(http.DefaultClient),
	}
}
