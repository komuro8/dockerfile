package main

import (
    "backend/database"
    "backend/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    r := gin.Default()

    // Usuarios
    r.POST("/usuarios", handlers.CrearUsuario)
    r.POST("/usuarios/iniciar-sesion", handlers.IniciarSesion)

    // Categorías
    r.POST("/categorias", handlers.CrearCategoria)
    r.GET("/categorias", handlers.ListarCategorias)
    r.DELETE("/categorias/:id", handlers.EliminarCategoria)

    // Tareas
    r.POST("/tareas", handlers.CrearTarea)
    r.GET("/tareas/usuario", handlers.ListarTareasUsuario)
    r.GET("/tareas/:id", handlers.ObtenerTareaPorID)
    r.PUT("/tareas/:id", handlers.ActualizarTarea)
    r.DELETE("/tareas/:id", handlers.EliminarTarea)

    r.Run(":8080")
}
