package dto

import "time"

type ComplaintCreateRequest struct {
	ComplaintNo     string `json:"complaint_no" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description"`
	Level           string `json:"level"`
	ComplainantName string `json:"complainant_name"`
	ComplainantTel  string `json:"complainant_tel"`
	EnterpriseName  string `json:"enterprise_name" binding:"required"`
	EnterpriseAddr  string `json:"enterprise_addr"`
	ResponsibleUser string `json:"responsible_user"`
	BatchNo         string `json:"batch_no"`
	Remark          string `json:"remark"`
}

type ComplaintUpdateRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Level           string `json:"level"`
	ComplainantName string `json:"complainant_name"`
	ComplainantTel  string `json:"complainant_tel"`
	EnterpriseName  string `json:"enterprise_name"`
	EnterpriseAddr  string `json:"enterprise_addr"`
	ResponsibleUser string `json:"responsible_user"`
	BatchNo         string `json:"batch_no"`
	Remark          string `json:"remark"`
}

type ComplaintTransitionRequest struct {
	ToStatus string `json:"to_status" binding:"required"`
	Reason   string `json:"reason"`
}

type ComplaintResponse struct {
	ID              uint      `json:"id"`
	ComplaintNo     string    `json:"complaint_no"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	Level           string    `json:"level"`
	ComplainantName string    `json:"complainant_name"`
	ComplainantTel  string    `json:"complainant_tel"`
	EnterpriseName  string    `json:"enterprise_name"`
	EnterpriseAddr  string    `json:"enterprise_addr"`
	ResponsibleUser string    `json:"responsible_user"`
	BatchNo         string    `json:"batch_no"`
	Priority        int       `json:"priority"`
	Remark          string    `json:"remark"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ComplaintListResponse struct {
	Total int                  `json:"total"`
	Items []ComplaintResponse `json:"items"`
}
