package dto

import "time"

type NoiseReadingCreateRequest struct {
	ReadingNo       string  `json:"reading_no" binding:"required"`
	SamplingPointID uint    `json:"sampling_point_id"`
	PointNo         string  `json:"point_no"`
	ComplaintID     uint    `json:"complaint_id"`
	ComplaintNo     string  `json:"complaint_no"`
	TimePeriodID    uint    `json:"time_period_id"`
	PeriodName      string  `json:"period_name"`
	MeasurementDate string  `json:"measurement_date" binding:"required"`
	MeasurementTime string  `json:"measurement_time" binding:"required"`
	Leq             float64 `json:"leq" binding:"required"`
	Lmax            float64 `json:"lmax"`
	Lmin            float64 `json:"lmin"`
	L10             float64 `json:"l10"`
	L90             float64 `json:"l90"`
	StandardLimit   float64 `json:"standard_limit"`
	ResponsibleUser string  `json:"responsible_user"`
	BatchNo         string  `json:"batch_no"`
	Remark          string  `json:"remark"`
}

type NoiseReadingUpdateRequest struct {
	PointNo         string  `json:"point_no"`
	PeriodName      string  `json:"period_name"`
	MeasurementDate string  `json:"measurement_date"`
	MeasurementTime string  `json:"measurement_time"`
	Leq             float64 `json:"leq"`
	Lmax            float64 `json:"lmax"`
	Lmin            float64 `json:"lmin"`
	L10             float64 `json:"l10"`
	L90             float64 `json:"l90"`
	StandardLimit   float64 `json:"standard_limit"`
	ResponsibleUser string  `json:"responsible_user"`
	BatchNo         string  `json:"batch_no"`
	Remark          string  `json:"remark"`
}

type NoiseReadingResponse struct {
	ID              uint      `json:"id"`
	ReadingNo       string    `json:"reading_no"`
	SamplingPointID uint      `json:"sampling_point_id"`
	PointNo         string    `json:"point_no"`
	ComplaintID     uint      `json:"complaint_id"`
	ComplaintNo     string    `json:"complaint_no"`
	TimePeriodID    uint      `json:"time_period_id"`
	PeriodName      string    `json:"period_name"`
	MeasurementDate time.Time `json:"measurement_date"`
	MeasurementTime string    `json:"measurement_time"`
	Leq             float64   `json:"leq"`
	Lmax            float64   `json:"lmax"`
	Lmin            float64   `json:"lmin"`
	L10             float64   `json:"l10"`
	L90             float64   `json:"l90"`
	StandardLimit   float64   `json:"standard_limit"`
	ExceedValue     float64   `json:"exceed_value"`
	IsExceeded      bool      `json:"is_exceeded"`
	Status          string    `json:"status"`
	ResponsibleUser string    `json:"responsible_user"`
	BatchNo         string    `json:"batch_no"`
	Remark          string    `json:"remark"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type NoiseReadingListResponse struct {
	Total int                    `json:"total"`
	Items []NoiseReadingResponse `json:"items"`
}
