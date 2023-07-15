package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-manage-api/internal/models"
	"server-manage-api/internal/services"
	"strings"
)

type Handler struct {
	serverManager services.ServerManagerService
}

func NewHandler(serverManager services.ServerManagerService) *Handler {
	return &Handler{
		serverManager: serverManager,
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}

func (h *Handler) GetServerByID(c *gin.Context) {
	serverID := c.Param("id")
	server, err := h.serverManager.GetByID(c, serverID)
	if server == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Server with id %s not found", serverID),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})

	}

	c.JSON(http.StatusOK, server)
}

func (h *Handler) DeleteServerByID(c *gin.Context) {
	serverID := c.Param("id")
	err := h.serverManager.DeleteByID(c, serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("Server with id %s was deleted", serverID)})
}

func (h *Handler) ListServers(c *gin.Context) {
	serversList, err := h.serverManager.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	if len(serversList) == 0 {
		c.JSON(http.StatusOK, gin.H{"total": len(serversList), "message": "there is no servers"})
		return
	}

	c.JSON(http.StatusOK, models.ListServerResponse{
		Servers: serversList,
		Total:   len(serversList),
	})
}

func (h *Handler) CreateServer(c *gin.Context) {
	var request models.CreateServerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "ShouldBindJSON : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	if err := validateRequest(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	created, err := h.serverManager.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, models.CreateServerResponse{
		Message: fmt.Sprintf("server %s was created with id %s", created.Name, created.ID),
		Status:  http.StatusOK,
	})
}

func validateRequest(request *models.CreateServerRequest) error {

	if request.Name == "" {
		return errors.New("name cannot be empty")
	}

	serverType := strings.ToLower(request.Type)
	switch serverType {
	case models.Small, models.Medium, models.Large:
		request.Type = serverType
	default:
		return errors.New("server type must be small, medium or large")
	}

	serverStatus := strings.ToLower(request.Status)
	switch serverStatus {
	case models.Starting, models.Running, models.Stopping, models.Stopped:
		request.Status = serverStatus
	default:
		return errors.New("server status, must be starting, running, stopping or stopped")
	}

	return nil
}
