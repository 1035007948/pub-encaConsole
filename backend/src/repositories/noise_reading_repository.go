package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type NoiseReadingRepository struct {
	db *gorm.DB
}

func NewNoiseReadingRepository() *NoiseReadingRepository {
	return &NoiseReadingRepository{db: database.GetDB()}
}

func (r *NoiseReadingRepository) Create(reading *models.NoiseReading) error {
	return r.db.Create(reading).Error
}

func (r *NoiseReadingRepository) Update(reading *models.NoiseReading) error {
	return r.db.Save(reading).Error
}

func (r *NoiseReadingRepository) Delete(id uint) error {
	return r.db.Delete(&models.NoiseReading{}, id).Error
}

func (r *NoiseReadingRepository) FindByID(id uint) (*models.NoiseReading, error) {
	var reading models.NoiseReading
	err := r.db.First(&reading, id).Error
	if err != nil {
		return nil, err
	}
	return &reading, nil
}

func (r *NoiseReadingRepository) FindByNo(readingNo string) (*models.NoiseReading, error) {
	var reading models.NoiseReading
	err := r.db.Where("reading_no = ?", readingNo).First(&reading).Error
	if err != nil {
		return nil, err
	}
	return &reading, nil
}

func (r *NoiseReadingRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.NoiseReading, int64, error) {
	var readings []models.NoiseReading
	var total int64

	query := r.db.Model(&models.NoiseReading{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if samplingPointID, ok := filters["sampling_point_id"]; ok {
		query = query.Where("sampling_point_id = ?", samplingPointID)
	}
	if complaintID, ok := filters["complaint_id"]; ok {
		query = query.Where("complaint_id = ?", complaintID)
	}
	if isExceeded, ok := filters["is_exceeded"]; ok {
		query = query.Where("is_exceeded = ?", isExceeded)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&readings).Error
	if err != nil {
		return nil, 0, err
	}

	return readings, total, nil
}

func (r *NoiseReadingRepository) FindBySamplingPointID(samplingPointID uint) ([]models.NoiseReading, error) {
	var readings []models.NoiseReading
	err := r.db.Where("sampling_point_id = ?", samplingPointID).Find(&readings).Error
	return readings, err
}

func (r *NoiseReadingRepository) FindByComplaintID(complaintID uint) ([]models.NoiseReading, error) {
	var readings []models.NoiseReading
	err := r.db.Where("complaint_id = ?", complaintID).Find(&readings).Error
	return readings, err
}

func (r *NoiseReadingRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.NoiseReading{}).Count(&count).Error
	return count, err
}

func (r *NoiseReadingRepository) CountExceeded() (int64, error) {
	var count int64
	err := r.db.Model(&models.NoiseReading{}).Where("is_exceeded = ?", true).Count(&count).Error
	return count, err
}

func (r *NoiseReadingRepository) AverageLeq() (float64, error) {
	var avg float64
	err := r.db.Model(&models.NoiseReading{}).Select("AVG(leq)").Scan(&avg).Error
	return avg, err
}
