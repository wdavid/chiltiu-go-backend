package handlers

import (
	"Proyectos_Go/internal/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(req.Name, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la categoría"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Categoría creada"})
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener categorías"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
