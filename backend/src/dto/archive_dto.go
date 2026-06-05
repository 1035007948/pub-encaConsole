package dto

import "time"

type ArchiveSnapshotCreateRequest struct {
	SnapshotName string `json:"snapshot_name" binding:"required"`
	SnapshotType string `json:"snapshot_type"`
	EntityType   string `json:"entity_type" binding:"required"`
	EntityIDs    []uint `json:"entity_ids" binding:"required"`
	Remark       string `json:"remark"`
}

type ArchiveSnapshotResponse struct {
	ID           uint      `json:"id"`
	SnapshotNo   string    `json:"snapshot_no"`
	SnapshotName string    `json:"snapshot_name"`
	SnapshotType string    `json:"snapshot_type"`
	EntityType   string    `json:"entity_type"`
	EntityIDs    string    `json:"entity_ids"`
	Data         string    `json:"data"`
	RecordCount  int       `json:"record_count"`
	Operator     string    `json:"operator"`
	BatchNo      string    `json:"batch_no"`
	Remark       string    `json:"remark"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ArchiveSnapshotListResponse struct {
	Total int                       `json:"total"`
	Items []ArchiveSnapshotResponse `json:"items"`
}

type ArchiveExportRequest struct {
	EntityType string `json:"entity_type" binding:"required"`
	EntityIDs  []uint `json:"entity_ids"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	BatchNo    string `json:"batch_no"`
	Format     string `json:"format"`
}

type ArchiveExportResponse struct {
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size"`
	RecordCount int    `json:"record_count"`
	DownloadURL string `json:"download_url"`
}
