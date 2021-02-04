package models

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Pagado         bool            `json:"is_pagado" gorm:"default:true"`
	Direccion      string          `json:"is_entregado" `
	Estado         bool            `json:"estado" gorm:"default:true"`
	ClienteID      int             `json:"id_cliente" `
	Cliente        Cliente         `json:"cliente"`
	DetallePedidos []DetallePedido `json:"detalles"`
}

func (Pedido) TableName() string {
	return "pedidos"
}
