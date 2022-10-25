package models

//go:generate go run github.com/dmarkham/enumer -type=ScanningStatus -json
type ScanningStatus int

const (
	NO_STATUS ScanningStatus = iota
	QUEUED
	IN_PROGRESS
	SUCCESS
	FAILURE
)
