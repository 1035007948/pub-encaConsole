package models

import (
	"time"

	"gorm.io/gorm"
)

type EvidenceAttachmentStatus string

const (
	EvidenceAttachmentStatusPending   EvidenceAttachmentStatus = "pending"
	EvidenceAttachmentStatusVerified  EvidenceAttachmentStatus = "verified"
	EvidenceAttachmentStatusRejected  EvidenceAttachmentStatus = "rejected"
	EvidenceAttachmentStatusArchived  EvidenceAttachmentStatus = "archived"
)

type EvidenceAttachment struct {
	ID              uint                     `gorm:"primaryKey" json:"id"`
	AttachmentNo    string                   `gorm:"uniqueIndex;size:50" json:"attachment_no"`
	FileName        string                   `gorm:"size:200" json:"file_name"`
	FilePath        string                   `gorm:"size:500" json:"file_path"`
	FileSize        int64                    `json:"file_size"`
	FileType        string                   `gorm:"size:50" json:"file_type"`
	MD5Hash         string                   `gorm:"size:32" json:"md5_hash"`
	ComplaintID     uint                     `json:"complaint_id"`
	ComplaintNo     string                   `gorm:"size:50" json:"complaint_no"`
	SamplingPointID uint                     `json:"sampling_point_id"`
	PointNo         string                   `gorm:"size:50" json:"point_no"`
	NoiseReadingID  uint                     `json:"noise_reading_id"`
	ReadingNo       string                   `gorm:"size:50" json:"reading_no"`
	AttachmentType  string                   `gorm:"size:50" json:"attachment_type"`
	Status          EvidenceAttachmentStatus `gorm:"size:20;default:pending" json:"status"`
	ResponsibleUser string                   `gorm:"size:100" json:"responsible_user"`
	BatchNo         string                   `gorm:"size:50" json:"batch_no"`
	Remark          string                   `gorm:"type:text" json:"remark"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	DeletedAt       gorm.DeletedAt           `gorm:"index" json:"-"`
}

func (EvidenceAttachment) TableName() string {
	return "evidence_attachments"
}
