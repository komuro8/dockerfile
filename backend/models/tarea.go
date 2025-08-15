package models

import "time"

type Tarea struct {
	ID                int       `json:"id"`
	Texto             string    `json:"texto"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaFinalizacion time.Time `json:"fecha_finalizacion"`
	Estado            string    `json:"estado"`
	IDCategoria       int       `json:"id_categoria"`
	IDUsuario         int       `json:"id_usuario"`
}
