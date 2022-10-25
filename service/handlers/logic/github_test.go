package logic

import (
	"encoding/base64"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/sonnht1409/scanning/service/models"
)

func TestServiceLogic_GetRepoContent(t *testing.T) {
	type args struct {
		owner string
		repo  string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.GithubFileContent
		wantErr bool
		mock    func(args args) *http.Client
	}{
		{
			name: "test all pass",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			want: []models.GithubFileContent{
				{
					Path:    "main.go",
					Url:     `https://github.com/owner1/repo1/main.go`,
					Content: base64.StdEncoding.EncodeToString([]byte("package main.go")),
				},
			},
			wantErr: false,
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				mockDefaultBranch := "master"

				mockSHA := base64.StdEncoding.EncodeToString([]byte("abc"))
				mockPath := "main.go"
				mockContent := base64.StdEncoding.EncodeToString([]byte("package main.go"))
				mockContentUrl := mockUrl + "/" + mockPath
				mockType := "blob"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposByOwnerByRepo,
						github.Repository{
							URL:           &mockUrl,
							DefaultBranch: &mockDefaultBranch,
						},
					),
					mock.WithRequestMatch(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						github.Tree{
							SHA: &mockDefaultBranch,
							Entries: []*github.TreeEntry{
								{
									SHA:  &mockSHA,
									Path: &mockPath,
									URL:  &mockContentUrl,
									Type: &mockType,
								},
							},
						},
					),
					mock.WithRequestMatch(
						mock.GetReposGitBlobsByOwnerByRepoByFileSha,
						github.Blob{
							Content: &mockContent,
						},
					),
				)
				return mockHttpClient
			},
		},
		{
			name: "test not get tree content",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			want:    nil,
			wantErr: false,
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				mockDefaultBranch := "master"

				mockSHA := base64.StdEncoding.EncodeToString([]byte("abc"))
				mockPath := "main"
				mockContentUrl := mockUrl + "/" + mockPath
				mockType := "tree"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposByOwnerByRepo,
						github.Repository{
							URL:           &mockUrl,
							DefaultBranch: &mockDefaultBranch,
						},
					),
					mock.WithRequestMatch(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						github.Tree{
							SHA: &mockDefaultBranch,
							Entries: []*github.TreeEntry{
								{
									SHA:  &mockSHA,
									Path: &mockPath,
									URL:  &mockContentUrl,
									Type: &mockType,
								},
							},
						},
					),
				)
				return mockHttpClient
			},
		},
		{
			name: "test get content failed",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			want:    nil,
			wantErr: false,
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				mockDefaultBranch := "master"

				mockSHA := base64.StdEncoding.EncodeToString([]byte("abc"))
				mockPath := "main.go"
				mockContentUrl := mockUrl + "/" + mockPath
				mockType := "blob"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposByOwnerByRepo,
						github.Repository{
							URL:           &mockUrl,
							DefaultBranch: &mockDefaultBranch,
						},
					),
					mock.WithRequestMatch(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						github.Tree{
							SHA: &mockDefaultBranch,
							Entries: []*github.TreeEntry{
								{
									SHA:  &mockSHA,
									Path: &mockPath,
									URL:  &mockContentUrl,
									Type: &mockType,
								},
							},
						},
					),
					mock.WithRequestMatchHandler(
						mock.GetReposGitBlobsByOwnerByRepoByFileSha,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
		},
		{
			name: "test get tree failed",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			want:    nil,
			wantErr: true,
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				mockDefaultBranch := "master"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposByOwnerByRepo,
						github.Repository{
							URL:           &mockUrl,
							DefaultBranch: &mockDefaultBranch,
						},
					),
					mock.WithRequestMatchHandler(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
		},
		{
			name: "test get repo failed",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			want:    nil,
			wantErr: true,
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatchHandler(
						mock.GetReposByOwnerByRepo,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceLogic()
			if tt.mock != nil {
				mockClient := tt.mock(tt.args)
				c := github.NewClient(mockClient)
				s.githubClient = c
			}
			got, err := s.GetRepoContent(tt.args.owner, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceLogic.GetRepoContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceLogic.GetRepoContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceLogic_GetRepoOwner(t *testing.T) {
	type fields struct {
		githubClient *github.Client
	}
	type args struct {
		repoUrl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test pass https owner + repo name",
			args: args{
				repoUrl: "https://github.com/owner/repo",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "owner",
		},
		{
			name: "test pass https owner",
			args: args{
				repoUrl: "https://github.com/owner",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "owner",
		},
		{
			name: "test pass http owner + repo name",
			args: args{
				repoUrl: "http://github.com/owner/repo1",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "owner",
		},
		{
			name: "test pass http owner",
			args: args{
				repoUrl: "http://github.com/owner",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "owner",
		},
		{
			name: "test fail",
			args: args{
				repoUrl: "http://github.com",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "",
		},
		{
			name: "test fail 1",
			args: args{
				repoUrl: "http://github.com/",
			},
			fields: fields{
				githubClient: github.NewClient(nil),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ServiceLogic{
				githubClient: tt.fields.githubClient,
			}
			if got := s.GetRepoOwner(tt.args.repoUrl); got != tt.want {
				t.Errorf("ServiceLogic.GetRepoOwner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceLogic_getFileContent(t *testing.T) {

	type args struct {
		owner string
		repo  string
		sha   string
	}
	mockContent := base64.StdEncoding.EncodeToString([]byte("package main.go"))
	tests := []struct {
		name    string
		args    args
		want    *github.Blob
		wantErr bool
		mock    func(args args) *http.Client
	}{
		{
			name: "test pass",
			args: args{
				owner: "owner1",
				repo:  "repo1",
				sha:   "sha1",
			},
			want:    nil,
			wantErr: true,
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatchHandler(
						mock.GetReposGitBlobsByOwnerByRepoByFileSha,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
		},
		{
			name: "test failed",
			args: args{
				owner: "owner1",
				repo:  "repo1",
				sha:   "sha1",
			},
			want: &github.Blob{
				Content: &mockContent,
			},
			wantErr: false,
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposGitBlobsByOwnerByRepoByFileSha,
						github.Blob{
							Content: &mockContent,
						},
					),
				)
				return mockHttpClient
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceLogic()
			if tt.mock != nil {
				mockClient := tt.mock(tt.args)
				c := github.NewClient(mockClient)
				s.githubClient = c
			}
			got, err := s.getFileContent(tt.args.owner, tt.args.repo, tt.args.sha)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceLogic.getFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceLogic.getFileContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceLogic_getAllFilePaths(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		branch string
	}
	wantTreeUrl := "https://github.com/owner/repo/main"
	wantBlobUrl := "https://github.com/owner/repo/main.go"
	mockSHA := base64.StdEncoding.EncodeToString([]byte("abc"))
	mockPath := "main.go"
	mockTreePath := "main"
	mockTypeBlob := "blob"
	mockTypeTree := "tree"
	tests := []struct {
		name    string
		args    args
		want    []*github.TreeEntry
		wantErr bool
		mock    func(args args) *http.Client
	}{
		{
			name: "test pass",
			args: args{
				owner:  "owner",
				repo:   "repo",
				branch: "master",
			},
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				// mockDefaultBranch := "master"

				mockContentUrl := mockUrl + "/" + mockPath
				mockType := "blob"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						github.Tree{
							Entries: []*github.TreeEntry{
								{
									SHA:  &mockSHA,
									Path: &mockPath,
									URL:  &mockContentUrl,
									Type: &mockType,
								},
							},
						},
					),
				)
				return mockHttpClient
			},
			want: []*github.TreeEntry{
				{
					SHA:  &mockSHA,
					Path: &mockPath,
					URL:  &wantBlobUrl,
					Type: &mockTypeBlob,
				},
			},
		},
		{
			name: "test pass with type tree",
			args: args{
				owner:  "owner",
				repo:   "repo",
				branch: "master",
			},
			mock: func(args args) *http.Client {
				mockUrl := "https://github.com/" + args.owner + "/" + args.repo
				// mockDefaultBranch := "master"

				mockContentUrl := mockUrl + "/" + mockTreePath
				mockType := "tree"
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						github.Tree{
							Entries: []*github.TreeEntry{
								{
									SHA:  &mockSHA,
									Path: &mockTreePath,
									URL:  &mockContentUrl,
									Type: &mockType,
								},
							},
						},
					),
				)
				return mockHttpClient
			},
			want: []*github.TreeEntry{
				{
					SHA:  &mockSHA,
					Path: &mockTreePath,
					URL:  &wantTreeUrl,
					Type: &mockTypeTree,
				},
			},
		},
		{
			name: "test failed",
			args: args{
				owner:  "owner",
				repo:   "repo",
				branch: "master",
			},
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatchHandler(
						mock.GetReposGitTreesByOwnerByRepoByTreeSha,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceLogic()
			if tt.mock != nil {
				mockClient := tt.mock(tt.args)
				c := github.NewClient(mockClient)
				s.githubClient = c
			}
			got, err := s.getAllFilePaths(tt.args.owner, tt.args.repo, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceLogic.getAllFilePaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceLogic.getAllFilePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceLogic_getRepository(t *testing.T) {
	type args struct {
		owner string
		repo  string
	}
	mockUrl := "https://github.com/owner1/repo1"
	mockDefaultBranch := "master"
	tests := []struct {
		name    string
		args    args
		want    *github.Repository
		wantErr bool
		mock    func(args args) *http.Client
	}{
		{
			name: "test pass",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatch(
						mock.GetReposByOwnerByRepo,
						github.Repository{
							URL:           &mockUrl,
							DefaultBranch: &mockDefaultBranch,
						},
					),
				)
				return mockHttpClient
			},
			want: &github.Repository{
				URL:           &mockUrl,
				DefaultBranch: &mockDefaultBranch,
			},
		},
		{
			name: "test failed",
			args: args{
				owner: "owner1",
				repo:  "repo1",
			},
			mock: func(args args) *http.Client {
				mockHttpClient := mock.NewMockedHTTPClient(
					mock.WithRequestMatchHandler(
						mock.GetReposByOwnerByRepo,
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							mock.WriteError(
								w,
								http.StatusInternalServerError,
								"github went belly up or something",
							)
						}),
					),
				)
				return mockHttpClient
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServiceLogic()
			if tt.mock != nil {
				mockClient := tt.mock(tt.args)
				c := github.NewClient(mockClient)
				s.githubClient = c
			}
			got, err := s.getRepository(tt.args.owner, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceLogic.getRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceLogic.getRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
