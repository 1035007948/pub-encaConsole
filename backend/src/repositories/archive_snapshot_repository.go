package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type ArchiveSnapshotRepository struct {
	db *gorm.DB
}

func NewArchiveSnapshotRepository() *ArchiveSnapshotRepository {
	return &ArchiveSnapshotRepository{db: database.GetDB()}
}

func (r *ArchiveSnapshotRepository) Create(snapshot *models.ArchiveSnapshot) error {
	return r.db.Create(snapshot).Error
}

func (r *ArchiveSnapshotRepository) FindByID(id uint) (*models.ArchiveSnapshot, error) {
	var snapshot models.ArchiveSnapshot
	err := r.db.First(&snapshot, id).Error
	if err != nil {
		return nil, err
	}
	return &snapshot, nil
}

func (r *ArchiveSnapshotRepository) FindAll(page, pageSize int) ([]models.ArchiveSnapshot, int64, error) {
	var snapshots []models.ArchiveSnapshot
	var total int64

	err := r.db.Model(&models.ArchiveSnapshot{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&snapshots).Error
	if err != nil {
		return nil, 0, err
	}

	return snapshots, total, nil
}
