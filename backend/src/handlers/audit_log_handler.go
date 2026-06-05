package handlers

import (
	"net/http"
	"strconv"

	"noise-complaint-backend/src/repositories"

	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	repo *repositories.AuditLogRepository
}

func NewAuditLogHandler() *AuditLogHandler {
	return &AuditLogHandler{repo: repositories.NewAuditLogRepository()}
}

func (h *AuditLogHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	filters := make(map[string]interface{})
	if action := c.Query("action"); action != "" {
		filters["action"] = action
	}
	if entityType := c.Query("entity_type"); entityType != "" {
		filters["entity_type"] = entityType
	}
	if operator := c.Query("operator"); operator != "" {
		filters["operator"] = operator
	}

	logs, total, err := h.repo.FindAll(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": logs,
	})
}
