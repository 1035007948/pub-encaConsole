package models

import (
	"time"

	"gorm.io/gorm"
)

type ArchiveSnapshot struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	SnapshotNo   string         `gorm:"uniqueIndex;size:50" json:"snapshot_no"`
	SnapshotName string         `gorm:"size:200" json:"snapshot_name"`
	SnapshotType string         `gorm:"size:50" json:"snapshot_type"`
	EntityType   string         `gorm:"size:50" json:"entity_type"`
	EntityIDs    string         `gorm:"type:text" json:"entity_ids"`
	Data         string         `gorm:"type:text" json:"data"`
	RecordCount  int            `json:"record_count"`
	Operator     string         `gorm:"size:100" json:"operator"`
	BatchNo      string         `gorm:"size:50" json:"batch_no"`
	Remark       string         `gorm:"type:text" json:"remark"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ArchiveSnapshot) TableName() string {
	return "archive_snapshots"
}
