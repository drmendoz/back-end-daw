package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	Nombre     string  `json:"nombre"`
	Precio     float64 `json:"precio"`
	Disponible bool    `json:"disponible" gorm:"default:true"`
	Estado     bool    `json:"estado" gorm:"default:true"`
}
