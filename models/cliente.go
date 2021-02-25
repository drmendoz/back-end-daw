package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Direccion string  `json:"direccion,omitempty"`
	UsuarioID int     `json:"id_usuario,omitempty"`
	Usuario   Usuario `json:"usuario,omitempty"`
	Estado    bool    `json:"estado" gorm:"default:true"`
}

func (Cliente) TableName() string {

	return "clientes"
}

type ClienteLog struct {
	Cliente Cliente `json:"respuesta"`
	Token   string  `json:"token"`
}
