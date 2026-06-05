package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type TimePeriodRepository struct {
	db *gorm.DB
}

func NewTimePeriodRepository() *TimePeriodRepository {
	return &TimePeriodRepository{db: database.GetDB()}
}

func (r *TimePeriodRepository) Create(period *models.TimePeriod) error {
	return r.db.Create(period).Error
}

func (r *TimePeriodRepository) BatchCreate(periods []models.TimePeriod) error {
	return r.db.Create(&periods).Error
}

func (r *TimePeriodRepository) Update(period *models.TimePeriod) error {
	return r.db.Save(period).Error
}

func (r *TimePeriodRepository) Delete(id uint) error {
	return r.db.Delete(&models.TimePeriod{}, id).Error
}

func (r *TimePeriodRepository) FindByID(id uint) (*models.TimePeriod, error) {
	var period models.TimePeriod
	err := r.db.First(&period, id).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (r *TimePeriodRepository) FindByNo(periodNo string) (*models.TimePeriod, error) {
	var period models.TimePeriod
	err := r.db.Where("period_no = ?", periodNo).First(&period).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}

func (r *TimePeriodRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.TimePeriod, int64, error) {
	var periods []models.TimePeriod
	var total int64

	query := r.db.Model(&models.TimePeriod{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if periodType, ok := filters["period_type"]; ok {
		query = query.Where("period_type = ?", periodType)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&periods).Error
	if err != nil {
		return nil, 0, err
	}

	return periods, total, nil
}

func (r *TimePeriodRepository) FindActive() ([]models.TimePeriod, error) {
	var periods []models.TimePeriod
	err := r.db.Where("status = ?", models.TimePeriodStatusActive).Find(&periods).Error
	return periods, err
}

func (r *TimePeriodRepository) FindByType(periodType models.TimePeriodType) (*models.TimePeriod, error) {
	var period models.TimePeriod
	err := r.db.Where("period_type = ? AND status = ?", periodType, models.TimePeriodStatusActive).First(&period).Error
	if err != nil {
		return nil, err
	}
	return &period, nil
}
