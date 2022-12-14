// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ScanningQuerySet

// ScanningQuerySet is an queryset type for Scanning
type ScanningQuerySet struct {
	db *gorm.DB
}

// NewScanningQuerySet constructs new ScanningQuerySet
func NewScanningQuerySet(db *gorm.DB) ScanningQuerySet {
	return ScanningQuerySet{
		db: db.Model(&Scanning{}),
	}
}

func (qs ScanningQuerySet) w(db *gorm.DB) ScanningQuerySet {
	return NewScanningQuerySet(db)
}

func (qs ScanningQuerySet) Select(fields ...ScanningDBSchemaField) ScanningQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Scanning) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Scanning) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) All(ret *[]Scanning) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) Count() (int64, error) {
	var count int64
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtEq(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtGt(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtGte(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtLt(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtLte(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) CreatedAtNe(createdAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) Delete() error {
	return qs.db.Delete(&Scanning{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(&Scanning{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(&Scanning{})
	return db.RowsAffected, db.Error
}

// FindingEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingEq(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding = ?", finding))
}

// FindingGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingGt(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding > ?", finding))
}

// FindingGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingGte(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding >= ?", finding))
}

// FindingIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingIn(finding ...string) ScanningQuerySet {
	if len(finding) == 0 {
		qs.db.AddError(errors.New("must at least pass one finding in FindingIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("finding IN (?)", finding))
}

// FindingIsNotNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingIsNotNull() ScanningQuerySet {
	return qs.w(qs.db.Where("finding IS NOT NULL"))
}

// FindingIsNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingIsNull() ScanningQuerySet {
	return qs.w(qs.db.Where("finding IS NULL"))
}

// FindingLike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingLike(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding LIKE ?", finding))
}

// FindingLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingLt(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding < ?", finding))
}

// FindingLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingLte(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding <= ?", finding))
}

// FindingNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingNe(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding != ?", finding))
}

// FindingNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingNotIn(finding ...string) ScanningQuerySet {
	if len(finding) == 0 {
		qs.db.AddError(errors.New("must at least pass one finding in FindingNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("finding NOT IN (?)", finding))
}

// FindingNotlike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FindingNotlike(finding string) ScanningQuerySet {
	return qs.w(qs.db.Where("finding NOT LIKE ?", finding))
}

// FinishedAtEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtEq(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at = ?", finishedAt))
}

// FinishedAtGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtGt(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at > ?", finishedAt))
}

// FinishedAtGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtGte(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at >= ?", finishedAt))
}

// FinishedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtIsNotNull() ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at IS NOT NULL"))
}

// FinishedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtIsNull() ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at IS NULL"))
}

// FinishedAtLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtLt(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at < ?", finishedAt))
}

// FinishedAtLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtLte(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at <= ?", finishedAt))
}

// FinishedAtNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) FinishedAtNe(finishedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("finished_at != ?", finishedAt))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) GetUpdater() ScanningUpdater {
	return NewScanningUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDEq(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDGt(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDGte(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDIn(ID ...int64) ScanningQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDLt(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDLte(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDNe(ID int64) ScanningQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) IDNotIn(ID ...int64) ScanningQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) Limit(limit int) ScanningQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) Offset(offset int) ScanningQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ScanningQuerySet) One(ret *Scanning) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByCreatedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByFinding is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByFinding() ScanningQuerySet {
	return qs.w(qs.db.Order("finding ASC"))
}

// OrderAscByFinishedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByFinishedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("finished_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByID() ScanningQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByQueuedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByQueuedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("queued_at ASC"))
}

// OrderAscByRepoName is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByRepoName() ScanningQuerySet {
	return qs.w(qs.db.Order("repo_name ASC"))
}

// OrderAscByRepoURL is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByRepoURL() ScanningQuerySet {
	return qs.w(qs.db.Order("repo_url ASC"))
}

// OrderAscByScanUniqueID is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByScanUniqueID() ScanningQuerySet {
	return qs.w(qs.db.Order("scan_unique_id ASC"))
}

// OrderAscByScanningAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByScanningAt() ScanningQuerySet {
	return qs.w(qs.db.Order("scanning_at ASC"))
}

// OrderAscByStatus is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByStatus() ScanningQuerySet {
	return qs.w(qs.db.Order("status ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderAscByUpdatedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByCreatedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByFinding is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByFinding() ScanningQuerySet {
	return qs.w(qs.db.Order("finding DESC"))
}

// OrderDescByFinishedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByFinishedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("finished_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByID() ScanningQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByQueuedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByQueuedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("queued_at DESC"))
}

// OrderDescByRepoName is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByRepoName() ScanningQuerySet {
	return qs.w(qs.db.Order("repo_name DESC"))
}

// OrderDescByRepoURL is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByRepoURL() ScanningQuerySet {
	return qs.w(qs.db.Order("repo_url DESC"))
}

// OrderDescByScanUniqueID is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByScanUniqueID() ScanningQuerySet {
	return qs.w(qs.db.Order("scan_unique_id DESC"))
}

// OrderDescByScanningAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByScanningAt() ScanningQuerySet {
	return qs.w(qs.db.Order("scanning_at DESC"))
}

// OrderDescByStatus is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByStatus() ScanningQuerySet {
	return qs.w(qs.db.Order("status DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) OrderDescByUpdatedAt() ScanningQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// QueuedAtEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtEq(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at = ?", queuedAt))
}

// QueuedAtGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtGt(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at > ?", queuedAt))
}

// QueuedAtGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtGte(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at >= ?", queuedAt))
}

// QueuedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtIsNotNull() ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at IS NOT NULL"))
}

// QueuedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtIsNull() ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at IS NULL"))
}

// QueuedAtLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtLt(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at < ?", queuedAt))
}

// QueuedAtLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtLte(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at <= ?", queuedAt))
}

// QueuedAtNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) QueuedAtNe(queuedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("queued_at != ?", queuedAt))
}

// RepoNameEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameEq(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name = ?", repoName))
}

// RepoNameGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameGt(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name > ?", repoName))
}

// RepoNameGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameGte(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name >= ?", repoName))
}

// RepoNameIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameIn(repoName ...string) ScanningQuerySet {
	if len(repoName) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoName in RepoNameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_name IN (?)", repoName))
}

// RepoNameLike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameLike(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name LIKE ?", repoName))
}

// RepoNameLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameLt(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name < ?", repoName))
}

// RepoNameLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameLte(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name <= ?", repoName))
}

// RepoNameNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameNe(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name != ?", repoName))
}

// RepoNameNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameNotIn(repoName ...string) ScanningQuerySet {
	if len(repoName) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoName in RepoNameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_name NOT IN (?)", repoName))
}

// RepoNameNotlike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoNameNotlike(repoName string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_name NOT LIKE ?", repoName))
}

// RepoURLEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLEq(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url = ?", repoURL))
}

// RepoURLGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLGt(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url > ?", repoURL))
}

// RepoURLGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLGte(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url >= ?", repoURL))
}

// RepoURLIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLIn(repoURL ...string) ScanningQuerySet {
	if len(repoURL) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoURL in RepoURLIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_url IN (?)", repoURL))
}

// RepoURLLike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLLike(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url LIKE ?", repoURL))
}

// RepoURLLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLLt(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url < ?", repoURL))
}

// RepoURLLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLLte(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url <= ?", repoURL))
}

// RepoURLNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLNe(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url != ?", repoURL))
}

// RepoURLNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLNotIn(repoURL ...string) ScanningQuerySet {
	if len(repoURL) == 0 {
		qs.db.AddError(errors.New("must at least pass one repoURL in RepoURLNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("repo_url NOT IN (?)", repoURL))
}

// RepoURLNotlike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) RepoURLNotlike(repoURL string) ScanningQuerySet {
	return qs.w(qs.db.Where("repo_url NOT LIKE ?", repoURL))
}

// ScanUniqueIDEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDEq(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id = ?", scanUniqueID))
}

// ScanUniqueIDGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDGt(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id > ?", scanUniqueID))
}

// ScanUniqueIDGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDGte(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id >= ?", scanUniqueID))
}

// ScanUniqueIDIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDIn(scanUniqueID ...string) ScanningQuerySet {
	if len(scanUniqueID) == 0 {
		qs.db.AddError(errors.New("must at least pass one scanUniqueID in ScanUniqueIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("scan_unique_id IN (?)", scanUniqueID))
}

// ScanUniqueIDLike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDLike(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id LIKE ?", scanUniqueID))
}

// ScanUniqueIDLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDLt(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id < ?", scanUniqueID))
}

// ScanUniqueIDLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDLte(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id <= ?", scanUniqueID))
}

// ScanUniqueIDNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDNe(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id != ?", scanUniqueID))
}

// ScanUniqueIDNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDNotIn(scanUniqueID ...string) ScanningQuerySet {
	if len(scanUniqueID) == 0 {
		qs.db.AddError(errors.New("must at least pass one scanUniqueID in ScanUniqueIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("scan_unique_id NOT IN (?)", scanUniqueID))
}

// ScanUniqueIDNotlike is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanUniqueIDNotlike(scanUniqueID string) ScanningQuerySet {
	return qs.w(qs.db.Where("scan_unique_id NOT LIKE ?", scanUniqueID))
}

// ScanningAtEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtEq(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at = ?", scanningAt))
}

// ScanningAtGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtGt(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at > ?", scanningAt))
}

// ScanningAtGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtGte(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at >= ?", scanningAt))
}

// ScanningAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtIsNotNull() ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at IS NOT NULL"))
}

// ScanningAtIsNull is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtIsNull() ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at IS NULL"))
}

// ScanningAtLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtLt(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at < ?", scanningAt))
}

// ScanningAtLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtLte(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at <= ?", scanningAt))
}

// ScanningAtNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) ScanningAtNe(scanningAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("scanning_at != ?", scanningAt))
}

// StatusEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusEq(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status = ?", status))
}

// StatusGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusGt(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status > ?", status))
}

// StatusGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusGte(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status >= ?", status))
}

// StatusIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusIn(status ...ScanningStatus) ScanningQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status IN (?)", status))
}

// StatusLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusLt(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status < ?", status))
}

// StatusLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusLte(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status <= ?", status))
}

// StatusNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusNe(status ScanningStatus) ScanningQuerySet {
	return qs.w(qs.db.Where("status != ?", status))
}

// StatusNotIn is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) StatusNotIn(status ...ScanningStatus) ScanningQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status NOT IN (?)", status))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtEq(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtGt(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtGte(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtLt(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtLte(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ScanningQuerySet) UpdatedAtNe(updatedAt time.Time) ScanningQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetCreatedAt(createdAt time.Time) ScanningUpdater {
	u.fields[string(ScanningDBSchema.CreatedAt)] = createdAt
	return u
}

// SetFinding is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetFinding(finding *string) ScanningUpdater {
	u.fields[string(ScanningDBSchema.Finding)] = finding
	return u
}

// SetFinishedAt is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetFinishedAt(finishedAt *time.Time) ScanningUpdater {
	u.fields[string(ScanningDBSchema.FinishedAt)] = finishedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetID(ID int64) ScanningUpdater {
	u.fields[string(ScanningDBSchema.ID)] = ID
	return u
}

// SetQueuedAt is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetQueuedAt(queuedAt *time.Time) ScanningUpdater {
	u.fields[string(ScanningDBSchema.QueuedAt)] = queuedAt
	return u
}

// SetRepoName is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetRepoName(repoName string) ScanningUpdater {
	u.fields[string(ScanningDBSchema.RepoName)] = repoName
	return u
}

// SetRepoURL is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetRepoURL(repoURL string) ScanningUpdater {
	u.fields[string(ScanningDBSchema.RepoURL)] = repoURL
	return u
}

// SetScanUniqueID is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetScanUniqueID(scanUniqueID string) ScanningUpdater {
	u.fields[string(ScanningDBSchema.ScanUniqueID)] = scanUniqueID
	return u
}

// SetScanningAt is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetScanningAt(scanningAt *time.Time) ScanningUpdater {
	u.fields[string(ScanningDBSchema.ScanningAt)] = scanningAt
	return u
}

// SetStatus is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetStatus(status ScanningStatus) ScanningUpdater {
	u.fields[string(ScanningDBSchema.Status)] = status
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) SetUpdatedAt(updatedAt time.Time) ScanningUpdater {
	u.fields[string(ScanningDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ScanningUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ScanningQuerySet

// ===== BEGIN of Scanning modifiers

// ScanningDBSchemaField describes database schema field. It requires for method 'Update'
type ScanningDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ScanningDBSchemaField) String() string {
	return string(f)
}

// ScanningDBSchema stores db field names of Scanning
var ScanningDBSchema = struct {
	ID           ScanningDBSchemaField
	RepoName     ScanningDBSchemaField
	RepoURL      ScanningDBSchemaField
	Finding      ScanningDBSchemaField
	ScanUniqueID ScanningDBSchemaField
	Status       ScanningDBSchemaField
	CreatedAt    ScanningDBSchemaField
	UpdatedAt    ScanningDBSchemaField
	QueuedAt     ScanningDBSchemaField
	FinishedAt   ScanningDBSchemaField
	ScanningAt   ScanningDBSchemaField
}{

	ID:           ScanningDBSchemaField("id"),
	RepoName:     ScanningDBSchemaField("repo_name"),
	RepoURL:      ScanningDBSchemaField("repo_url"),
	Finding:      ScanningDBSchemaField("finding"),
	ScanUniqueID: ScanningDBSchemaField("scan_unique_id"),
	Status:       ScanningDBSchemaField("status"),
	CreatedAt:    ScanningDBSchemaField("created_at"),
	UpdatedAt:    ScanningDBSchemaField("updated_at"),
	QueuedAt:     ScanningDBSchemaField("queued_at"),
	FinishedAt:   ScanningDBSchemaField("finished_at"),
	ScanningAt:   ScanningDBSchemaField("scanning_at"),
}

// Update updates Scanning fields by primary key
// nolint: dupl
func (o *Scanning) Update(db *gorm.DB, fields ...ScanningDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":             o.ID,
		"repo_name":      o.RepoName,
		"repo_url":       o.RepoURL,
		"finding":        o.Finding,
		"scan_unique_id": o.ScanUniqueID,
		"status":         o.Status,
		"created_at":     o.CreatedAt,
		"updated_at":     o.UpdatedAt,
		"queued_at":      o.QueuedAt,
		"finished_at":    o.FinishedAt,
		"scanning_at":    o.ScanningAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Scanning %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ScanningUpdater is an Scanning updates manager
type ScanningUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewScanningUpdater creates new Scanning updater
// nolint: dupl
func NewScanningUpdater(db *gorm.DB) ScanningUpdater {
	return ScanningUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Scanning{}),
	}
}

// ===== END of Scanning modifiers

// ===== END of all query sets
