package handlers

import (
	"net/http"

	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/seed"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *gin.Context) {
	db := database.GetDB()
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"message": "database not connected",
		})
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"message": "failed to get database connection",
		})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"message": "database ping failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "all systems operational",
	})
}

type SeedHandler struct{}

func NewSeedHandler() *SeedHandler {
	return &SeedHandler{}
}

func (h *SeedHandler) Reset(c *gin.Context) {
	db := database.GetDB()

	db.Exec("TRUNCATE TABLE complaints CASCADE")
	db.Exec("TRUNCATE TABLE sampling_points CASCADE")
	db.Exec("TRUNCATE TABLE noise_readings CASCADE")
	db.Exec("TRUNCATE TABLE time_periods CASCADE")
	db.Exec("TRUNCATE TABLE evidence_attachments CASCADE")
	db.Exec("TRUNCATE TABLE rectification_measures CASCADE")
	db.Exec("TRUNCATE TABLE retest_records CASCADE")
	db.Exec("TRUNCATE TABLE status_transitions CASCADE")
	db.Exec("TRUNCATE TABLE rule_configs CASCADE")
	db.Exec("TRUNCATE TABLE anomaly_events CASCADE")
	db.Exec("TRUNCATE TABLE audit_logs CASCADE")
	db.Exec("TRUNCATE TABLE archive_snapshots CASCADE")

	seed.SeedAll()

	c.JSON(http.StatusOK, gin.H{
		"message": "seed data reset successfully",
	})
}

func (h *SeedHandler) Browse(c *gin.Context) {
	db := database.GetDB()

	var complaints []models.Complaint
	var samplingPoints []models.SamplingPoint
	var noiseReadings []models.NoiseReading
	var timePeriods []models.TimePeriod

	db.Limit(10).Find(&complaints)
	db.Limit(10).Find(&samplingPoints)
	db.Limit(10).Find(&noiseReadings)
	db.Limit(10).Find(&timePeriods)

	c.JSON(http.StatusOK, gin.H{
		"complaints": gin.H{
			"total": len(complaints),
			"items": complaints,
		},
		"sampling_points": gin.H{
			"total": len(samplingPoints),
			"items": samplingPoints,
		},
		"noise_readings": gin.H{
			"total": len(noiseReadings),
			"items": noiseReadings,
		},
		"time_periods": gin.H{
			"total": len(timePeriods),
			"items": timePeriods,
		},
	})
}
