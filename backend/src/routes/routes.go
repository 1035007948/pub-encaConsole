package routes

import (
	"noise-complaint-backend/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		complaintHandler := handlers.NewComplaintHandler()
		complaints := api.Group("/complaints")
		{
			complaints.GET("", complaintHandler.GetList)
			complaints.POST("", complaintHandler.Create)
			complaints.GET("/:id", complaintHandler.GetByID)
			complaints.PUT("/:id", complaintHandler.Update)
			complaints.DELETE("/:id", complaintHandler.Delete)
			complaints.POST("/:id/transition", complaintHandler.Transition)
		}

		samplingPointHandler := handlers.NewSamplingPointHandler()
		samplingPoints := api.Group("/sampling-points")
		{
			samplingPoints.GET("", samplingPointHandler.GetList)
			samplingPoints.POST("", samplingPointHandler.Create)
			samplingPoints.GET("/:id", samplingPointHandler.GetByID)
			samplingPoints.PUT("/:id", samplingPointHandler.Update)
			samplingPoints.DELETE("/:id", samplingPointHandler.Delete)
			samplingPoints.POST("/:id/transition", samplingPointHandler.Transition)
			samplingPoints.GET("/:id/validate", samplingPointHandler.ValidateConsistency)
		}

		noiseReadingHandler := handlers.NewNoiseReadingHandler()
		noiseReadings := api.Group("/noise-readings")
		{
			noiseReadings.GET("", noiseReadingHandler.GetList)
			noiseReadings.POST("", noiseReadingHandler.Create)
			noiseReadings.GET("/:id", noiseReadingHandler.GetByID)
			noiseReadings.PUT("/:id", noiseReadingHandler.Update)
			noiseReadings.DELETE("/:id", noiseReadingHandler.Delete)
			noiseReadings.GET("/by-point/:id", noiseReadingHandler.GetBySamplingPointID)
		}

		timePeriodHandler := handlers.NewTimePeriodHandler()
		timePeriods := api.Group("/time-periods")
		{
			timePeriods.GET("", timePeriodHandler.GetList)
			timePeriods.POST("", timePeriodHandler.Create)
			timePeriods.GET("/:id", timePeriodHandler.GetByID)
			timePeriods.PUT("/:id", timePeriodHandler.Update)
			timePeriods.DELETE("/:id", timePeriodHandler.Delete)
			timePeriods.POST("/batch-import", timePeriodHandler.BatchImport)
		}

		anomalyHandler := handlers.NewAnomalyHandler()
		anomalies := api.Group("/anomalies")
		{
			anomalies.GET("", anomalyHandler.GetList)
			anomalies.GET("/:id", anomalyHandler.GetByID)
			anomalies.PUT("/:id/resolve", anomalyHandler.Resolve)
		}

		calculationHandler := handlers.NewCalculationHandler()
		calculate := api.Group("/calculate")
		{
			calculate.POST("/priority", calculationHandler.CalculatePriority)
			calculate.POST("/completeness", calculationHandler.CalculateCompleteness)
			calculate.POST("/compliance", calculationHandler.CheckCompliance)
		}

		statisticsHandler := handlers.NewStatisticsHandler()
		statistics := api.Group("/statistics")
		{
			statistics.GET("/dashboard", statisticsHandler.GetDashboard)
			statistics.GET("/completeness", statisticsHandler.GetCompletenessStatistics)
			statistics.GET("/rectification-rate", statisticsHandler.GetRectificationRateStatistics)
			statistics.GET("/retest-pass-rate", statisticsHandler.GetRetestPassRateStatistics)
		}

		archiveHandler := handlers.NewArchiveHandler()
		archive := api.Group("/archive")
		{
			archive.POST("/snapshot", archiveHandler.CreateSnapshot)
			archive.GET("/snapshots", archiveHandler.GetSnapshotList)
			archive.POST("/export", archiveHandler.Export)
		}

		auditLogHandler := handlers.NewAuditLogHandler()
		api.GET("/audit-logs", auditLogHandler.GetList)

		ruleConfigHandler := handlers.NewRuleConfigHandler()
		rules := api.Group("/rules")
		{
			rules.GET("", ruleConfigHandler.GetList)
			rules.POST("", ruleConfigHandler.Create)
			rules.GET("/:id", ruleConfigHandler.GetByID)
			rules.PUT("/:id", ruleConfigHandler.Update)
		}

		healthHandler := handlers.NewHealthHandler()
		api.GET("/health", healthHandler.Check)

		seedHandler := handlers.NewSeedHandler()
		seed := api.Group("/seed")
		{
			seed.POST("/reset", seedHandler.Reset)
			seed.GET("/browse", seedHandler.Browse)
		}
	}
}
