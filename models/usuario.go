package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre     string `json:"nombres,omitempty" `
	Correo     string `json:"correo,omitempty" `
	Rol        string `json:"rol,omitempty"`
	Contrasena string `json:"contrasena,omitempty" `
	Estado     bool   `json:"estado" gorm:"default:true"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
