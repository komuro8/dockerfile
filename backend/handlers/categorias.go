package handlers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearCategoria(c *gin.Context) {
	var categoria models.Categoria
	if err := c.ShouldBindJSON(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec(
		"INSERT INTO categorias (nombre, descripcion) VALUES ($1, $2)",
		categoria.Nombre, categoria.Descripcion,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Categoría creada"})
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
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM categorias WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada"})
}
