package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type AnomalyEventRepository struct {
	db *gorm.DB
}

func NewAnomalyEventRepository() *AnomalyEventRepository {
	return &AnomalyEventRepository{db: database.GetDB()}
}

func (r *AnomalyEventRepository) Create(event *models.AnomalyEvent) error {
	return r.db.Create(event).Error
}

func (r *AnomalyEventRepository) Update(event *models.AnomalyEvent) error {
	return r.db.Save(event).Error
}

func (r *AnomalyEventRepository) Delete(id uint) error {
	return r.db.Delete(&models.AnomalyEvent{}, id).Error
}

func (r *AnomalyEventRepository) FindByID(id uint) (*models.AnomalyEvent, error) {
	var event models.AnomalyEvent
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *AnomalyEventRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.AnomalyEvent, int64, error) {
	var events []models.AnomalyEvent
	var total int64

	query := r.db.Model(&models.AnomalyEvent{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if severity, ok := filters["severity"]; ok {
		query = query.Where("severity = ?", severity)
	}
	if eventType, ok := filters["event_type"]; ok {
		query = query.Where("event_type = ?", eventType)
	}
	if entityType, ok := filters["entity_type"]; ok {
		query = query.Where("entity_type = ?", entityType)
	}
	if entityID, ok := filters["entity_id"]; ok {
		query = query.Where("entity_id = ?", entityID)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&events).Error
	if err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

func (r *AnomalyEventRepository) FindOpen() ([]models.AnomalyEvent, error) {
	var events []models.AnomalyEvent
	err := r.db.Where("status IN ?", []models.AnomalyEventStatus{
		models.AnomalyEventStatusOpen,
		models.AnomalyEventStatusProcessing,
	}).Find(&events).Error
	return events, err
}

func (r *AnomalyEventRepository) CountOpen() (int64, error) {
	var count int64
	err := r.db.Model(&models.AnomalyEvent{}).Where("status IN ?", []models.AnomalyEventStatus{
		models.AnomalyEventStatusOpen,
		models.AnomalyEventStatusProcessing,
	}).Count(&count).Error
	return count, err
}

func (r *AnomalyEventRepository) Resolve(id uint, resolutionNote string) error {
	return r.db.Model(&models.AnomalyEvent{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":         models.AnomalyEventStatusResolved,
		"resolution_note": resolutionNote,
	}).Error
}
