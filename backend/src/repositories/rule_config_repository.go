package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type RuleConfigRepository struct {
	db *gorm.DB
}

func NewRuleConfigRepository() *RuleConfigRepository {
	return &RuleConfigRepository{db: database.GetDB()}
}

func (r *RuleConfigRepository) Create(rule *models.RuleConfig) error {
	return r.db.Create(rule).Error
}

func (r *RuleConfigRepository) Update(rule *models.RuleConfig) error {
	return r.db.Save(rule).Error
}

func (r *RuleConfigRepository) Delete(id uint) error {
	return r.db.Delete(&models.RuleConfig{}, id).Error
}

func (r *RuleConfigRepository) FindByID(id uint) (*models.RuleConfig, error) {
	var rule models.RuleConfig
	err := r.db.First(&rule, id).Error
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *RuleConfigRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.RuleConfig, int64, error) {
	var rules []models.RuleConfig
	var total int64

	query := r.db.Model(&models.RuleConfig{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if ruleType, ok := filters["rule_type"]; ok {
		query = query.Where("rule_type = ?", ruleType)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("priority desc, created_at desc").Find(&rules).Error
	if err != nil {
		return nil, 0, err
	}

	return rules, total, nil
}

func (r *RuleConfigRepository) FindActive() ([]models.RuleConfig, error) {
	var rules []models.RuleConfig
	err := r.db.Where("status = ?", models.RuleConfigStatusActive).
		Order("priority desc").
		Find(&rules).Error
	return rules, err
}
