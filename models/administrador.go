package models

import "gorm.io/gorm"

type Administrador struct {
	gorm.Model
	Nombre     string `json:"nombres,omitempty" binding:"required"`
	Correo     string `json:"correo,omitempty" binding:"required"`
	Rol        string `json:"rol,omitempty"`
	Contrasena string `json:"contrasena,omitempty" binding:"required"`
	Imagen     string `json:"imagen,omitempty"`
	Estado     string `json:"estado,omitempty"`
}

func (Administrador) TableName() string {
	return "administradores"
}
