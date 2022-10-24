package models

//go:generate go run github.com/dmarkham/enumer -type=ScanningStatus -json
type ScanningStatus int

const (
	QUEUED ScanningStatus = iota
	IN_PROGRESS
	SUCCESS
	FAILURE
)
