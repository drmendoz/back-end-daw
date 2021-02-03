package models

import (
	"time"

	"gorm.io/gorm"
)

type Afiliado struct {
	gorm.Model
	Nombre             string    `json:"nombres,omitempty" binding:"required"`
	Tipo               string    `json:"tipo,omitempty" binding:"required"`
	Descripcion        string    `json:"descripcion,omitempty"`
	Imagen             string    `json:"imagen,omitempty"`
	Estado             bool      `json:"estado" gorm:"default:true"`
	Correo             string    `json:"correo,omitempty" binding:"required"`
	Clave              string    `json:"clave,omitempty" binding:"required"`
	ImagenPresentacion string    `json:"imagen_presentacion,omitempty"`
	ImagenBackground   string    `json:"imagen_background,omitempty"`
	CategoriaID        int       `json:"id_categoria,omitempty"`
	Categoria          Categoria `json:"categoria" gorm:"embed"`
	HoraEntrada        time.Time `json:"hora_entrada"`
	HoraSalida         time.Time `json:"hora_salida"`
	Referencia         string    `json:"referencia,omitempty"`
}

type InsertAfiliado struct {
	gorm.Model
	Nombre             string    `json:"nombres,omitempty" binding:"required"`
	Tipo               string    `json:"tipo,omitempty" binding:"required"`
	Descripcion        string    `json:"descripcion,omitempty"`
	Imagen             string    `json:"imagen,omitempty"`
	Estado             bool      `json:"estado,omitempty" gorm:"default:true"`
	Correo             string    `json:"correo,omitempty" binding:"required"`
	Clave              string    `json:"clave,omitempty" binding:"required"`
	ImagenPresentacion string    `json:"imagen_presentacion,omitempty"`
	ImagenBackground   string    `json:"imagen_background,omitempty"`
	CategoriaID        int       `json:"id_categoria,omitempty"`
	HoraEntrada        time.Time `json:"hora_entrada"`
	HoraSalida         time.Time `json:"hora_salida"`
	Referencia         string    `json:"referencia,omitempty"`
}

func (InsertAfiliado) TableName() string {
	return "afiliados"

}

func (Afiliado) TableName() string {
	return "afiliados"
}
