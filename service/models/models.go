package models

import (
	"encoding/json"
	"time"

	"github.com/sonnht1409/scanning/service/common"
	"gorm.io/gorm"
)

//go:generate goqueryset -in models.go

// gen:qs
type Scanning struct {
	ID           int64          `json:"id,omitempty"`
	RepoName     string         `json:"repo_name,omitempty"`
	RepoURL      string         `json:"repo_url,omitempty"`
	Finding      *string        `json:"finding,omitempty"`
	ScanUniqueID string         `json:"scan_unique_id,omitempty"`
	Status       ScanningStatus `json:"scanning_status"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	QueuedAt     *time.Time     `json:"queued_at,omitempty"`
	FinishedAt   *time.Time     `json:"finished_at,omitempty"`
	ScanningAt   *time.Time     `json:"scanning_at,omitempty"`
	Findings     []Finding      `json:"findings,omitempty" gorm:"-"`
}

func (s *Scanning) TableName() string {
	return "scannings"
}

func (s *Scanning) EncodeFindings() string {
	if len(s.Findings) > 0 {
		findingStr := common.Stringify(s.Findings)
		return findingStr
	}
	return ""
}

func (s *Scanning) AfterFind(db *gorm.DB) error {
	if s.Finding != nil {
		//nolint: errcheck
		_ = json.Unmarshal([]byte(*s.Finding), &s.Findings)
	}
	return nil
}

func (s *Scanning) AppendFindings(f Finding) {
	if len(s.Findings) > 0 {
		s.Findings = append(s.Findings, f)
		return
	}
	s.Findings = []Finding{f}
}

func (s *Scanning) SighUpScanning(db *gorm.DB) error {
	return db.Create(&s).Error
}

func (s *Scanning) GetByScanID(db *gorm.DB, scanId string) error {
	return NewScanningQuerySet(db).ScanUniqueIDEq(scanId).One(s)
}

func (s *Scanning) UpdateScanningStatus(db *gorm.DB, status ScanningStatus) error {
	query := NewScanningQuerySet(db).IDEq(s.ID).GetUpdater().SetStatus(status)
	if status == IN_PROGRESS {
		scanningAt := time.Now()
		query = query.SetScanningAt(&scanningAt)
	}
	if status == FAILURE || status == SUCCESS {
		finishedAt := time.Now()
		query = query.SetFinishedAt(&finishedAt)
	}
	return query.Update()
}

func (s *Scanning) UpdateFindings(db *gorm.DB) error {
	if len(s.Findings) == 0 {
		return nil
	}
	findingStr := s.EncodeFindings()
	return NewScanningQuerySet(db).IDEq(s.ID).GetUpdater().SetFinding(&findingStr).Update()
}

func (s *Scanning) ViewScanResultByRepoName(db *gorm.DB) ([]Scanning, error) {
	result := []Scanning{}
	err := NewScanningQuerySet(db).RepoNameEq(s.RepoName).OrderDescByID().All(&result)
	return result, err
}

func (s *Scanning) ViewOneScanningProcess(db *gorm.DB) (bool, error) {
	err := NewScanningQuerySet(db).ScanUniqueIDEq(s.ScanUniqueID).One(s)
	if err == nil {
		return true, nil
	}

	if err == gorm.ErrRecordNotFound {
		return false, err
	}

	return true, err
}
