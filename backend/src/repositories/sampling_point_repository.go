package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type SamplingPointRepository struct {
	db *gorm.DB
}

func NewSamplingPointRepository() *SamplingPointRepository {
	return &SamplingPointRepository{db: database.GetDB()}
}

func (r *SamplingPointRepository) Create(point *models.SamplingPoint) error {
	return r.db.Create(point).Error
}

func (r *SamplingPointRepository) Update(point *models.SamplingPoint) error {
	return r.db.Save(point).Error
}

func (r *SamplingPointRepository) Delete(id uint) error {
	return r.db.Delete(&models.SamplingPoint{}, id).Error
}

func (r *SamplingPointRepository) FindByID(id uint) (*models.SamplingPoint, error) {
	var point models.SamplingPoint
	err := r.db.First(&point, id).Error
	if err != nil {
		return nil, err
	}
	return &point, nil
}

func (r *SamplingPointRepository) FindByNo(pointNo string) (*models.SamplingPoint, error) {
	var point models.SamplingPoint
	err := r.db.Where("point_no = ?", pointNo).First(&point).Error
	if err != nil {
		return nil, err
	}
	return &point, nil
}

func (r *SamplingPointRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.SamplingPoint, int64, error) {
	var points []models.SamplingPoint
	var total int64

	query := r.db.Model(&models.SamplingPoint{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if complaintID, ok := filters["complaint_id"]; ok {
		query = query.Where("complaint_id = ?", complaintID)
	}
	if responsibleUser, ok := filters["responsible_user"]; ok {
		query = query.Where("responsible_user = ?", responsibleUser)
	}
	if batchNo, ok := filters["batch_no"]; ok {
		query = query.Where("batch_no = ?", batchNo)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&points).Error
	if err != nil {
		return nil, 0, err
	}

	return points, total, nil
}

func (r *SamplingPointRepository) FindByComplaintID(complaintID uint) ([]models.SamplingPoint, error) {
	var points []models.SamplingPoint
	err := r.db.Where("complaint_id = ?", complaintID).Find(&points).Error
	return points, err
}

func (r *SamplingPointRepository) UpdateStatus(id uint, status models.SamplingPointStatus) error {
	return r.db.Model(&models.SamplingPoint{}).Where("id = ?", id).Update("status", status).Error
}

func (r *SamplingPointRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.SamplingPoint{}).Count(&count).Error
	return count, err
}
