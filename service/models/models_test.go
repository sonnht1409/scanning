package models

import (
	"database/sql/driver"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

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

func TestScanning_GetByScanID(t *testing.T) {
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
		db     *gorm.DB
		scanId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "test pass",
			args: args{
				scanId: uuid.NewString(),
				db:     gormDB,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "repo_name", "repo_url", "finding",
					"scan_unique_id", "status", "created_at", "updated_at", "queued_at", "scanning_at", "finished_at"}).
					AddRow(1, "repo1", "repoUrl1", nil, args.scanId, QUEUED, time.Now(), time.Now(), nil, nil, nil)

				query := regexp.QuoteMeta("SELECT * FROM `scannings` WHERE scan_unique_id = ? ORDER BY `scannings`.`id` LIMIT 1")
				mock.ExpectQuery(query).WithArgs(args.scanId).
					WillReturnRows(rows)
			},
		},
		{
			name: "test failed",
			args: args{
				scanId: uuid.NewString(),
				db:     gormDB,
			},
			mock: func(args args) {

				query := regexp.QuoteMeta("SELECT * FROM `scannings` WHERE scan_unique_id = ? ORDER BY `scannings`.`id` LIMIT 1")
				mock.ExpectQuery(query).WithArgs(args.scanId).WillReturnError(errors.New("not found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanning{
				ScanUniqueID: tt.args.scanId,
			}
			if tt.mock != nil {
				tt.mock(tt.args)
			}
			if err := s.GetByScanID(tt.args.db, tt.args.scanId); (err != nil) != tt.wantErr {
				t.Errorf("Scanning.GetByScanID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}

func TestScanning_UpdateScanningStatus(t *testing.T) {
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

	// mockTime := time.Now()
	type args struct {
		db     *gorm.DB
		status ScanningStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args args, fields fields)
	}{
		{
			name: "test pass scanning at",
			mock: func(args args, fields fields) {
				query := "UPDATE `scannings` SET `scanning_at`=?,`status`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(AnyTime{}, args.status, AnyTime{}, fields.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				db:     gormDB,
				status: IN_PROGRESS,
			},
			fields: fields{
				ID: 1,
			},
		},
		{
			name: "test pass success finished at",
			mock: func(args args, fields fields) {
				query := "UPDATE `scannings` SET `finished_at`=?,`status`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(AnyTime{}, args.status, AnyTime{}, fields.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				db:     gormDB,
				status: SUCCESS,
			},
			fields: fields{
				ID: 1,
			},
		},
		{
			name: "test pass failure finished at",
			mock: func(args args, fields fields) {
				query := "UPDATE `scannings` SET `finished_at`=?,`status`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(AnyTime{}, args.status, AnyTime{}, fields.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				db:     gormDB,
				status: FAILURE,
			},
			fields: fields{
				ID: 1,
			},
		},
		{
			name: "test fail update status",
			mock: func(args args, fields fields) {
				query := "UPDATE `scannings` SET `scanning_at`=?,`status`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(AnyTime{}, args.status, AnyTime{}, fields.ID).WillReturnError(errors.New("error"))
			},
			args: args{
				db:     gormDB,
				status: IN_PROGRESS,
			},
			fields: fields{
				ID: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(tt.args, tt.fields)
			}
			s := &Scanning{
				ID:         tt.fields.ID,
				Status:     tt.fields.Status,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
				ScanningAt: tt.fields.ScanningAt,
				QueuedAt:   tt.fields.QueuedAt,
				FinishedAt: tt.fields.FinishedAt,
			}
			if err := s.UpdateScanningStatus(tt.args.db, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Scanning.UpdateScanningStatus() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}

func TestScanning_UpdateFindings(t *testing.T) {
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
		mock    func(fields)
	}{
		{
			name: "test pass",
			args: args{
				db: gormDB,
			},
			fields: fields{
				ID: 1,
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
			mock: func(fields fields) {
				query := "UPDATE `scannings` SET `finding`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(fields.Finding, AnyTime{}, fields.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "test failed",
			args: args{
				db: gormDB,
			},
			fields: fields{
				ID: 1,
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
			mock: func(fields fields) {
				query := "UPDATE `scannings` SET `finding`=?,`updated_at`=? WHERE id = ?"
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(fields.Finding, AnyTime{}, fields.ID).WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanning{
				ID:       tt.fields.ID,
				Findings: tt.fields.Findings,
			}
			finding := s.EncodeFindings()
			tt.fields.Finding = &finding
			if tt.mock != nil {
				tt.mock(tt.fields)
			}
			if err := s.UpdateFindings(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Scanning.UpdateFindings() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}

func TestScanning_ViewScanResultByRepoName(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{SkipDefaultTransaction: true})
	type fields struct {
		RepoName string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Scanning
		wantErr bool
		mock    func(fields fields)
	}{
		{
			name: "test pass",
			args: args{
				db: gormDB,
			},
			fields: fields{
				RepoName: "repo_name",
			},
			mock: func(fields fields) {
				query := "SELECT * FROM `scannings` WHERE repo_name = ? ORDER BY id DESC"
				rows := sqlmock.NewRows([]string{"id", "repo_name", "repo_url", "status"}).AddRow(1, "repo_name", "repo_url", QUEUED)
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(fields.RepoName).WillReturnRows(rows)
			},
			want: []Scanning{
				{
					ID:       1,
					RepoName: "repo_name",
					RepoURL:  "repo_url",
					Status:   QUEUED,
				},
			},
		},
		{
			name: "test fail",
			args: args{
				db: gormDB,
			},
			fields: fields{
				RepoName: "repo_name",
			},
			mock: func(fields fields) {
				query := "SELECT * FROM `scannings` WHERE repo_name = ? ORDER BY id DESC"
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(fields.RepoName).WillReturnError(errors.New("error"))
			},
			want:    []Scanning{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanning{
				RepoName: tt.fields.RepoName,
			}
			if tt.mock != nil {
				tt.mock(tt.fields)
			}
			got, err := s.ViewScanResultByRepoName(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanning.ViewScanResultByRepoName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanning.ViewScanResultByRepoName() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}

func TestScanning_ViewOneScanningProcess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{SkipDefaultTransaction: true})
	type fields struct {
		ScanUniqueID string
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		mock    func(fields fields)
	}{
		{
			name: "test pass",
			args: args{
				db: gormDB,
			},
			fields: fields{
				ScanUniqueID: uuid.NewString(),
			},
			want: true,
			mock: func(fields fields) {
				rows := sqlmock.NewRows([]string{"id", "repo_name", "repo_url", "finding",
					"scan_unique_id", "status", "created_at", "updated_at", "queued_at", "scanning_at", "finished_at"}).
					AddRow(1, "repo1", "repoUrl1", nil, fields.ScanUniqueID, QUEUED, time.Now(), time.Now(), nil, nil, nil)

				query := regexp.QuoteMeta("SELECT * FROM `scannings` WHERE scan_unique_id = ? ORDER BY `scannings`.`id` LIMIT 1")
				mock.ExpectQuery(query).WithArgs(fields.ScanUniqueID).
					WillReturnRows(rows)
			},
		},
		{
			name: "test fail",
			args: args{
				db: gormDB,
			},
			fields: fields{
				ScanUniqueID: uuid.NewString(),
			},
			want:    true,
			wantErr: true,
			mock: func(fields fields) {

				query := regexp.QuoteMeta("SELECT * FROM `scannings` WHERE scan_unique_id = ? ORDER BY `scannings`.`id` LIMIT 1")
				mock.ExpectQuery(query).WithArgs(fields.ScanUniqueID).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanning{
				ScanUniqueID: tt.fields.ScanUniqueID,
			}
			if tt.mock != nil {
				tt.mock(tt.fields)
			}
			got, err := s.ViewOneScanningProcess(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanning.ViewOneScanningProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Scanning.ViewOneScanningProcess() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Scanning.ExpectationsWereMet() error = %v", err)
			}
		})
	}
}
