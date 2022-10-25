package models

import "time"

type GithubFileContent struct {
	Path    string `json:"path"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

type RegexRule struct {
	RuleName  string `json:"rule_name"`
	RuleValue string `json:"rule_value"`
}

type Location struct {
	Path string `json:"path"`
	Line int    `json:"line"`
}

type Finding struct {
	RuleName string   `json:"rule_name"`
	Location Location `json:"location"`
}

type CreateScanningRequest struct {
	RepoName string `json:"repo_name,omitempty"`
	RepoURL  string `json:"repo_url,omitempty"`
}

type CreateScanningResponse struct {
	ID           int    `json:"id"`
	ScanUniqueID string `json:"scan_unique_id,omitempty"`
}

type Signal struct {
	ScanID string `json:"scan_id"`
	Status string `json:"status"`
}

type ViewResultRequest struct {
	RepoName string `json:"repo_name" form:"repo_name"`
}

type ViewResultResponse struct {
	ID           int64      `json:"id,omitempty"`
	RepoName     string     `json:"repo_name,omitempty"`
	RepoURL      string     `json:"repo_url,omitempty"`
	ScanUniqueID string     `json:"scan_unique_id,omitempty"`
	Status       string     `json:"scanning_status"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	QueuedAt     *time.Time `json:"queued_at,omitempty"`
	FinishedAt   *time.Time `json:"finished_at,omitempty"`
	ScanningAt   *time.Time `json:"scanning_at,omitempty"`
	Findings     []Finding  `json:"findings,omitempty"`
}

type ViewScanningProcessRequest struct {
	ScanUniqueID string `json:"scan_unique_id,omitempty" form:"scan_unique_id"`
}

type RetryScanProcessRequest struct {
	ScanUniqueID string `json:"scan_unique_id,omitempty" form:"scan_unique_id"`
}
