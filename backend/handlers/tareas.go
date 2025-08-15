package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearTarea(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Crear tarea"})
}

func ListarTareasUsuario(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Listar tareas por usuario"})
}

func ObtenerTareaPorID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Obtener tarea por ID"})
}

func ActualizarTarea(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Actualizar tarea"})
}

func EliminarTarea(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Eliminar tarea"})
}
