package handlers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearCategoria(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Crear categoría"})
}

func ListarCategorias(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, nombre, descripcion FROM categorias")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var cat models.Categoria
		if err := rows.Scan(&cat.ID, &cat.Nombre, &cat.Descripcion); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categorias = append(categorias, cat)
	}

	c.JSON(http.StatusOK, categorias)
}

func EliminarCategoria(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Eliminar categoría"})
}
