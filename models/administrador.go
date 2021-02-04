package models

import "gorm.io/gorm"

type Administrador struct {
	gorm.Model
	UsuarioID int     `json:"id_usuario"`
	Usuario   Usuario `json:"usuario"`
	Estado    bool    `json:"estado" gorm:"default:true"`
}

func (Administrador) TableName() string {
	return "administradores"
}
