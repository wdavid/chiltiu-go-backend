package handlers

import (
	"Proyectos_Go/internal/core/model"
	"Proyectos_Go/internal/core/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TourismHandler struct {
	service *service.TourismService
}

func NewTourismHandler(s *service.TourismService) *TourismHandler {
	return &TourismHandler{service: s}
}

func (h *TourismHandler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	destinations, total, err := h.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener destinos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": destinations,
		"meta": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total_items":  total,
			"total_pages":  (total + int64(limit) - 1) / int64(limit),
		},
	})
}

func (h *TourismHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	dest, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destino no encontrado"})
		return
	}
	c.JSON(http.StatusOK, dest)
}

func (h *TourismHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("description")
	loc := c.PostForm("location")
	latStr := c.PostForm("latitude")
	lonStr := c.PostForm("longitude")
	catIDStr := c.PostForm("category_id")
	catID, _ := strconv.Atoi(catIDStr)

	imageFile, _ := c.FormFile("image")
	videoFile, _ := c.FormFile("video")

	lat, _ := strconv.ParseFloat(latStr, 64)
	lon, _ := strconv.ParseFloat(lonStr, 64)

	dest := &model.TouristDestination{
		Name:        name,
		Description: desc,
		Location:    loc,
		Latitude:    lat,
		Longitude:   lon,
		CategoryID:  uint(catID),
	}

	if err := h.service.CreateDestination(dest, imageFile, videoFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fullDest, err := h.service.GetByID(dest.ID)

	responseDest := dest
	if err == nil {
		responseDest = fullDest
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Destino creado", "data": responseDest})
}

func (h *TourismHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	desc := c.PostForm("description")
	loc := c.PostForm("location")
	lat, _ := strconv.ParseFloat(c.PostForm("latitude"), 64)
	lon, _ := strconv.ParseFloat(c.PostForm("longitude"), 64)
	catID, _ := strconv.Atoi(c.PostForm("category_id"))

	imageFile, _ := c.FormFile("image")
	videoFile, _ := c.FormFile("video")

	updateData := &model.TouristDestination{
		Name:        name,
		Description: desc,
		Location:    loc,
		Latitude:    lat,
		Longitude:   lon,
		CategoryID:  uint(catID),
	}

	updatedDest, err := h.service.UpdateDestination(uint(id), updateData, imageFile, videoFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Destino actualizado", "data": updatedDest})
}

func (h *TourismHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteDestination(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Destino eliminado correctamente"})
}
