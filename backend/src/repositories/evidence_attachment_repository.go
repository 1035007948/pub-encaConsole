package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type EvidenceAttachmentRepository struct {
	db *gorm.DB
}

func NewEvidenceAttachmentRepository() *EvidenceAttachmentRepository {
	return &EvidenceAttachmentRepository{db: database.GetDB()}
}

func (r *EvidenceAttachmentRepository) Create(attachment *models.EvidenceAttachment) error {
	return r.db.Create(attachment).Error
}

func (r *EvidenceAttachmentRepository) Update(attachment *models.EvidenceAttachment) error {
	return r.db.Save(attachment).Error
}

func (r *EvidenceAttachmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.EvidenceAttachment{}, id).Error
}

func (r *EvidenceAttachmentRepository) FindByID(id uint) (*models.EvidenceAttachment, error) {
	var attachment models.EvidenceAttachment
	err := r.db.First(&attachment, id).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *EvidenceAttachmentRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.EvidenceAttachment, int64, error) {
	var attachments []models.EvidenceAttachment
	var total int64

	query := r.db.Model(&models.EvidenceAttachment{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if complaintID, ok := filters["complaint_id"]; ok {
		query = query.Where("complaint_id = ?", complaintID)
	}
	if samplingPointID, ok := filters["sampling_point_id"]; ok {
		query = query.Where("sampling_point_id = ?", samplingPointID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&attachments).Error
	if err != nil {
		return nil, 0, err
	}

	return attachments, total, nil
}

func (r *EvidenceAttachmentRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.EvidenceAttachment{}).Count(&count).Error
	return count, err
}

func (r *EvidenceAttachmentRepository) CountByComplaintID(complaintID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.EvidenceAttachment{}).Where("complaint_id = ?", complaintID).Count(&count).Error
	return count, err
}
