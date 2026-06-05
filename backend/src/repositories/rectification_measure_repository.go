package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type RectificationMeasureRepository struct {
	db *gorm.DB
}

func NewRectificationMeasureRepository() *RectificationMeasureRepository {
	return &RectificationMeasureRepository{db: database.GetDB()}
}

func (r *RectificationMeasureRepository) Create(measure *models.RectificationMeasure) error {
	return r.db.Create(measure).Error
}

func (r *RectificationMeasureRepository) Update(measure *models.RectificationMeasure) error {
	return r.db.Save(measure).Error
}

func (r *RectificationMeasureRepository) FindByID(id uint) (*models.RectificationMeasure, error) {
	var measure models.RectificationMeasure
	err := r.db.First(&measure, id).Error
	if err != nil {
		return nil, err
	}
	return &measure, nil
}

func (r *RectificationMeasureRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.RectificationMeasure, int64, error) {
	var measures []models.RectificationMeasure
	var total int64

	query := r.db.Model(&models.RectificationMeasure{})

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
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&measures).Error
	if err != nil {
		return nil, 0, err
	}

	return measures, total, nil
}

func (r *RectificationMeasureRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.RectificationMeasure{}).Count(&count).Error
	return count, err
}

func (r *RectificationMeasureRepository) CountByStatus(status models.RectificationMeasureStatus) (int64, error) {
	var count int64
	err := r.db.Model(&models.RectificationMeasure{}).Where("status = ?", status).Count(&count).Error
	return count, err
}
