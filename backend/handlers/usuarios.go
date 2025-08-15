package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearUsuario(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Crear usuario"})
}

func IniciarSesion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: Iniciar sesión"})
}
