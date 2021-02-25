package models

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	Pagado         bool            `json:"is_pagado" gorm:"default:true"`
	Entregado      bool            `json:"is_entregado" `
	Direccion      string          `json:"direccion"`
	Estado         bool            `json:"estado" gorm:"default:true"`
	ClienteID      int             `json:"id_cliente" `
	Cliente        Cliente         `json:"cliente"`
	DetallePedidos []DetallePedido `json:"detalles"  `
	Subtotal       float64         `json:"subtotal"`
	CostoEnvio     float64         `json:"costo_envio"`
	Total          float64         `json:"total"`
	Ciudad         string          `json:"ciudad"`
}

func (Pedido) TableName() string {
	return "pedidos"
}
