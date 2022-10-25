// Code generated by mockery v2.14.0. DO NOT EDIT.

package logic

import (
	github "github.com/google/go-github/v48/github"
	mock "github.com/stretchr/testify/mock"

	models "github.com/sonnht1409/scanning/service/models"
)

// MockIServiceLogic is an autogenerated mock type for the IServiceLogic type
type MockIServiceLogic struct {
	mock.Mock
}

// CheckRule provides a mock function with given fields: content, rule
func (_m *MockIServiceLogic) CheckRule(content string, rule models.RegexRule) (bool, []int) {
	ret := _m.Called(content, rule)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, models.RegexRule) bool); ok {
		r0 = rf(content, rule)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []int
	if rf, ok := ret.Get(1).(func(string, models.RegexRule) []int); ok {
		r1 = rf(content, rule)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]int)
		}
	}

	return r0, r1
}

// GetRepoContent provides a mock function with given fields: _a0, _a1
func (_m *MockIServiceLogic) GetRepoContent(_a0 string, _a1 string) ([]models.GithubFileContent, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []models.GithubFileContent
	if rf, ok := ret.Get(0).(func(string, string) []models.GithubFileContent); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.GithubFileContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepoOwner provides a mock function with given fields: repoUrl
func (_m *MockIServiceLogic) GetRepoOwner(repoUrl string) string {
	ret := _m.Called(repoUrl)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(repoUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getAllFilePaths provides a mock function with given fields: owner, repo, branch
func (_m *MockIServiceLogic) getAllFilePaths(owner string, repo string, branch string) ([]*github.TreeEntry, error) {
	ret := _m.Called(owner, repo, branch)

	var r0 []*github.TreeEntry
	if rf, ok := ret.Get(0).(func(string, string, string) []*github.TreeEntry); ok {
		r0 = rf(owner, repo, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.TreeEntry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repo, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getFileContent provides a mock function with given fields: owner, repo, sha
func (_m *MockIServiceLogic) getFileContent(owner string, repo string, sha string) (*github.Blob, error) {
	ret := _m.Called(owner, repo, sha)

	var r0 *github.Blob
	if rf, ok := ret.Get(0).(func(string, string, string) *github.Blob); ok {
		r0 = rf(owner, repo, sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repo, sha)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getRepository provides a mock function with given fields: owner, repo
func (_m *MockIServiceLogic) getRepository(owner string, repo string) (*github.Repository, error) {
	ret := _m.Called(owner, repo)

	var r0 *github.Repository
	if rf, ok := ret.Get(0).(func(string, string) *github.Repository); ok {
		r0 = rf(owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockIServiceLogic interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIServiceLogic creates a new instance of MockIServiceLogic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIServiceLogic(t mockConstructorTestingTNewMockIServiceLogic) *MockIServiceLogic {
	mock := &MockIServiceLogic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}