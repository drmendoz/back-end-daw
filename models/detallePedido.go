package models

import "gorm.io/gorm"

type DetallePedido struct {
	gorm.Model
	ProductoID int     `json:"id_producto"`
	Cantidad   int     `json:"cantidad"`
	Precio     float64 `json:"precio"`
	Subtotal   float64 `json:"subtotal"`
	CostoEnvio float64 `json:"costo_envio"`
	Total      float64 `json:"total"`
	PedidoID   int     `json:"id_pedido"`
	Estado     bool    `json:"estado" gorm:"default:true"`
}

func (DetallePedido) TableName() string {
	return "detalle_pedidos"
}
