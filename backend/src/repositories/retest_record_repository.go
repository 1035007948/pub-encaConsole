package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type RetestRecordRepository struct {
	db *gorm.DB
}

func NewRetestRecordRepository() *RetestRecordRepository {
	return &RetestRecordRepository{db: database.GetDB()}
}

func (r *RetestRecordRepository) Create(record *models.RetestRecord) error {
	return r.db.Create(record).Error
}

func (r *RetestRecordRepository) Update(record *models.RetestRecord) error {
	return r.db.Save(record).Error
}

func (r *RetestRecordRepository) FindByID(id uint) (*models.RetestRecord, error) {
	var record models.RetestRecord
	err := r.db.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RetestRecordRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.RetestRecord, int64, error) {
	var records []models.RetestRecord
	var total int64

	query := r.db.Model(&models.RetestRecord{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if complaintID, ok := filters["complaint_id"]; ok {
		query = query.Where("complaint_id = ?", complaintID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

func (r *RetestRecordRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.RetestRecord{}).Count(&count).Error
	return count, err
}

func (r *RetestRecordRepository) CountPassed() (int64, error) {
	var count int64
	err := r.db.Model(&models.RetestRecord{}).Where("is_passed = ?", true).Count(&count).Error
	return count, err
}
