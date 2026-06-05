package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type StatusTransitionRepository struct {
	db *gorm.DB
}

func NewStatusTransitionRepository() *StatusTransitionRepository {
	return &StatusTransitionRepository{db: database.GetDB()}
}

func (r *StatusTransitionRepository) Create(transition *models.StatusTransition) error {
	return r.db.Create(transition).Error
}

func (r *StatusTransitionRepository) FindByEntity(entityType string, entityID uint) ([]models.StatusTransition, error) {
	var transitions []models.StatusTransition
	err := r.db.Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Order("created_at desc").
		Find(&transitions).Error
	return transitions, err
}

func (r *StatusTransitionRepository) FindAll(page, pageSize int) ([]models.StatusTransition, int64, error) {
	var transitions []models.StatusTransition
	var total int64

	err := r.db.Model(&models.StatusTransition{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&transitions).Error
	if err != nil {
		return nil, 0, err
	}

	return transitions, total, nil
}
