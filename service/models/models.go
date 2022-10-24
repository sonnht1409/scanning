package models

import (
	"time"
)

//go:generate goqueryset -in models.go

// gen:qs
type Scanning struct {
	ID           int64          `json:"id,omitempty"`
	RepoName     string         `json:"repo_name,omitempty"`
	RepoURL      string         `json:"repo_url,omitempty"`
	Finding      string         `json:"finding,omitempty"`
	ScanUniqueID string         `json:"scan_unique_id,omitempty"`
	Status       ScanningStatus `json:"scanning_status"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	QueuedAt     time.Time      `json:"queued_at,omitempty"`
	FinishedAt   time.Time      `json:"finished_at,omitempty"`
	ScanningAt   time.Time      `json:"scanning_at,omitempty"`
}
