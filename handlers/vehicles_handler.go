package handlers

import (
	"net/http"
	"service-golang/models"
	"service-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehiclesHandler struct {
	VehiclesService *services.VehiclesService
}

func NewVehcilesHandler(service *services.VehiclesService) *VehiclesHandler {
	return &VehiclesHandler{VehiclesService: service}
}

func (h *VehiclesHandler) CreateVehicles(c *gin.Context) {
	var vehicles models.Vehicles

	if err := c.ShouldBind(&vehicles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VehiclesService.CreateVehicles(&vehicles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Create Vehicles"})
}

func (h *VehiclesHandler) GetAllVehicles(c *gin.Context) {
	vehicles, err := h.VehiclesService.GetAllVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": vehicles, "message": "OK"})
}

func (h *VehiclesHandler) UpdateVehicles(c *gin.Context) {
	idParam := c.Param("id")
	var vehicles models.Vehicles

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&vehicles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VehiclesService.UpdateVehicles(&vehicles, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Update Vehicles"})

}

func (h *VehiclesHandler) DeleteVehicles(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VehiclesService.DeleteVehicles(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Delete Data Vehicles"})
}
