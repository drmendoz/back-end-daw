package models

import "gorm.io/gorm"

type Resena struct {
	gorm.Model
	Comentario string  `json:"comentario"`
	ClienteID  int     `json:"id_cliente"`
	Cliente    Cliente `json:"cliente,omitempty"`
	Ciudad     string  `json:"ciudad"`
	Estado     bool    `json:"estado" gorm:"default:true"`
}

func (Resena) TableName() string {
	return "resenas"
}
