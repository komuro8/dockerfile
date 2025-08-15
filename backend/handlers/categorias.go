package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearCategoria(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Crear categoría"})
}

func ListarCategorias(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Listar categorías"})
}

func EliminarCategoria(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Eliminar categoría"})
}
