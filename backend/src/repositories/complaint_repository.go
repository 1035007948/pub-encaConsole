package repositories

import (
	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

type ComplaintRepository struct {
	db *gorm.DB
}

func NewComplaintRepository() *ComplaintRepository {
	return &ComplaintRepository{db: database.GetDB()}
}

func (r *ComplaintRepository) Create(complaint *models.Complaint) error {
	return r.db.Create(complaint).Error
}

func (r *ComplaintRepository) Update(complaint *models.Complaint) error {
	return r.db.Save(complaint).Error
}

func (r *ComplaintRepository) Delete(id uint) error {
	return r.db.Delete(&models.Complaint{}, id).Error
}

func (r *ComplaintRepository) FindByID(id uint) (*models.Complaint, error) {
	var complaint models.Complaint
	err := r.db.First(&complaint, id).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func (r *ComplaintRepository) FindByNo(complaintNo string) (*models.Complaint, error) {
	var complaint models.Complaint
	err := r.db.Where("complaint_no = ?", complaintNo).First(&complaint).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func (r *ComplaintRepository) FindAll(page, pageSize int, filters map[string]interface{}) ([]models.Complaint, int64, error) {
	var complaints []models.Complaint
	var total int64

	query := r.db.Model(&models.Complaint{})

	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if level, ok := filters["level"]; ok {
		query = query.Where("level = ?", level)
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
	err = query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&complaints).Error
	if err != nil {
		return nil, 0, err
	}

	return complaints, total, nil
}

func (r *ComplaintRepository) UpdateStatus(id uint, status models.ComplaintStatus) error {
	return r.db.Model(&models.Complaint{}).Where("id = ?", id).Update("status", status).Error
}

func (r *ComplaintRepository) UpdatePriority(id uint, priority int) error {
	return r.db.Model(&models.Complaint{}).Where("id = ?", id).Update("priority", priority).Error
}

func (r *ComplaintRepository) CountByStatus(status models.ComplaintStatus) (int64, error) {
	var count int64
	err := r.db.Model(&models.Complaint{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (r *ComplaintRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Complaint{}).Count(&count).Error
	return count, err
}
