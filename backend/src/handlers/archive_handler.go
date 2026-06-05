package handlers

import (
	"net/http"
	"strconv"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/services"

	"github.com/gin-gonic/gin"
)

type ArchiveHandler struct {
	service *services.ArchiveService
}

func NewArchiveHandler() *ArchiveHandler {
	return &ArchiveHandler{service: services.NewArchiveService()}
}

func (h *ArchiveHandler) CreateSnapshot(c *gin.Context) {
	var req dto.ArchiveSnapshotCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operator := c.GetHeader("X-Operator")
	if operator == "" {
		operator = "system"
	}

	snapshot, err := h.service.CreateSnapshot(&req, operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, snapshot)
}

func (h *ArchiveHandler) GetSnapshotList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	response, err := h.service.GetSnapshotList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ArchiveHandler) Export(c *gin.Context) {
	var req dto.ArchiveExportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.Export(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
