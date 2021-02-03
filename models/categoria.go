package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Nombre string `json:"nombre" binding:"required"`
	Estado bool   `json:"estado" gorm:"default:true"`
	Imagen string `json:"imagen"`
}

func (Categoria) TableName() string {
	return "categorias"
}
