package dto

import "time"

type SamplingPointCreateRequest struct {
	PointNo           string  `json:"point_no" binding:"required"`
	PointName         string  `json:"point_name" binding:"required"`
	Address           string  `json:"address"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	ComplaintID       uint    `json:"complaint_id"`
	ComplaintNo       string  `json:"complaint_no"`
	ResponsibleUser   string  `json:"responsible_user"`
	ScheduledDate     string  `json:"scheduled_date"`
	ScheduledTimeFrom string  `json:"scheduled_time_from"`
	ScheduledTimeTo   string  `json:"scheduled_time_to"`
	BatchNo           string  `json:"batch_no"`
	Remark            string  `json:"remark"`
}

type SamplingPointUpdateRequest struct {
	PointName         string  `json:"point_name"`
	Address           string  `json:"address"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	ResponsibleUser   string  `json:"responsible_user"`
	ScheduledDate     string  `json:"scheduled_date"`
	ScheduledTimeFrom string  `json:"scheduled_time_from"`
	ScheduledTimeTo   string  `json:"scheduled_time_to"`
	BatchNo           string  `json:"batch_no"`
	Remark            string  `json:"remark"`
}

type SamplingPointTransitionRequest struct {
	ToStatus string `json:"to_status" binding:"required"`
	Reason   string `json:"reason"`
}

type SamplingPointResponse struct {
	ID                uint       `json:"id"`
	PointNo           string     `json:"point_no"`
	PointName         string     `json:"point_name"`
	Address           string     `json:"address"`
	Longitude         float64    `json:"longitude"`
	Latitude          float64    `json:"latitude"`
	Status            string     `json:"status"`
	ComplaintID       uint       `json:"complaint_id"`
	ComplaintNo       string     `json:"complaint_no"`
	ResponsibleUser   string     `json:"responsible_user"`
	ScheduledDate     *time.Time `json:"scheduled_date"`
	ScheduledTimeFrom string     `json:"scheduled_time_from"`
	ScheduledTimeTo   string     `json:"scheduled_time_to"`
	BatchNo           string     `json:"batch_no"`
	Remark            string     `json:"remark"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type SamplingPointListResponse struct {
	Total int                     `json:"total"`
	Items []SamplingPointResponse `json:"items"`
}
