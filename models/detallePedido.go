package models

import "gorm.io/gorm"

type DetallePedido struct {
	gorm.Model
	ProductoID int      `json:"id_producto"`
	Producto   Producto `json:"producto"`
	Cantidad   int      `json:"cantidad"`
	Precio     float64  `json:"precio"`
	PedidoID   int      `json:"id_pedido"`
	Estado     bool     `json:"estado" gorm:"default:true"`
}

func (DetallePedido) TableName() string {
	return "detalle_pedidos"
}
