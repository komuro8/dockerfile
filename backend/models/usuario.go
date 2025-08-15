package models

type Usuario struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre_usuario"`
	Contraseña   string `json:"contraseña"`
	ImagenPerfil string `json:"imagen_perfil"`
}
