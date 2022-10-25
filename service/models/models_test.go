package models

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestScanning_EncodeFindings(t *testing.T) {

	type fields struct {
		ID           int64
		RepoName     string
		RepoURL      string
		Finding      *string
		ScanUniqueID string
		Status       ScanningStatus
		CreatedAt    time.Time
		UpdatedAt    time.Time
		QueuedAt     *time.Time
		FinishedAt   *time.Time
		ScanningAt   *time.Time
		Findings     []Finding
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		mock   func(fields fields, mockSQL sqlmock.Sqlmock)
	}{
		{
			name: "test pass",
			fields: fields{
				Findings: []Finding{
					{
						Location: Location{
							Path: "main.go",
							Line: 1,
						},
						RuleName: "PublicKeyCheck",
					},
				},
			},
			want: `[{"rule_name":"PublicKeyCheck","location":{"path":"main.go","line":1}}]`,
		},
		{
			name: "test nil",
			fields: fields{
				Findings: nil,
			},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanning{
				ID:           tt.fields.ID,
				RepoName:     tt.fields.RepoName,
				RepoURL:      tt.fields.RepoURL,
				Finding:      tt.fields.Finding,
				ScanUniqueID: tt.fields.ScanUniqueID,
				Status:       tt.fields.Status,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				QueuedAt:     tt.fields.QueuedAt,
				FinishedAt:   tt.fields.FinishedAt,
				ScanningAt:   tt.fields.ScanningAt,
				Findings:     tt.fields.Findings,
			}
			if got := s.EncodeFindings(); got != tt.want {
				t.Errorf("Scanning.EncodeFindings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanning_SighUpScanning(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{SkipDefaultTransaction: true})
	type fields struct {
		ID           int64
		RepoName     string
		RepoURL      string
		Finding      *string
		ScanUniqueID string
		Status       ScanningStatus
		CreatedAt    time.Time
		UpdatedAt    time.Time
		QueuedAt     *time.Time
		FinishedAt   *time.Time
		ScanningAt   *time.Time
		Findings     []Finding
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(fields fields)
	}{
		{
			name: "test pass",
			mock: func(fields fields) {
				mock.ExpectExec("INSERT INTO `scannings`").WithArgs(fields.RepoName, fields.RepoURL, fields.Finding,
					fields.ScanUniqueID, fields.Status, fields.CreatedAt, fields.UpdatedAt,
					fields.QueuedAt, fields.FinishedAt, fields.FinishedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			fields: fields{
				RepoName:     "repo_name",
				RepoURL:      "repo_url",
				ScanUniqueID: uuid.NewString(),
				Status:       QUEUED,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			args: args{
				db: gormDB,
			},
		},
		{
			name: "test failed",
			mock: func(fields fields) {
				mock.ExpectExec("INSERT INTO `scannings`").WithArgs(fields.RepoName, fields.RepoURL, fields.Finding,
					fields.ScanUniqueID, fields.Status, fields.CreatedAt, fields.UpdatedAt,
					fields.QueuedAt, fields.FinishedAt, fields.FinishedAt).
					WillReturnError(errors.New("some error"))
			},
			fields: fields{
				RepoName:     "repo_name",
				RepoURL:      "repo_url",
				ScanUniqueID: uuid.NewString(),
				Status:       QUEUED,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: true,
			args: args{
				db: gormDB,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.db = gormDB
			s := &Scanning{
				ID:           tt.fields.ID,
				RepoName:     tt.fields.RepoName,
				RepoURL:      tt.fields.RepoURL,
				Finding:      tt.fields.Finding,
				ScanUniqueID: tt.fields.ScanUniqueID,
				Status:       tt.fields.Status,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				QueuedAt:     tt.fields.QueuedAt,
				FinishedAt:   tt.fields.FinishedAt,
				ScanningAt:   tt.fields.ScanningAt,
				Findings:     tt.fields.Findings,
			}
			if tt.mock != nil {
				tt.mock(tt.fields)
			}
			if err := s.SighUpScanning(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Scanning.SighUpScanning() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}
