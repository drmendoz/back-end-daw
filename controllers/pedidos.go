package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
)

func GetPedidos(c *gin.Context) {
	pedidos := []*models.Pedido{}
	err := models.Db.Where("pedidos.estado = ?", true).Preload("Cliente.Usuario").Preload("DetallePedidos.Producto").Preload(clause.Associations).Find(&pedidos).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener pedidos"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, pedidos, c, http.StatusOK)
}

func CreatePedido(c *gin.Context) {
	pedido := &models.Pedido{}
	err := c.ShouldBindJSON(pedido)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	result := models.Db.Create(pedido)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear pedido"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, pedido, c, http.StatusCreated)

}

func UpdatePedido(c *gin.Context) {

	pedido := &models.Pedido{}

	err := c.ShouldBindJSON(pedido)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	ui, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	pedido.ID = uint(ui)
	result := models.Db.Updates(pedido)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar pedido"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, pedido, c, http.StatusOK)

}

func GetPedidoPorId(c *gin.Context) {
	pedido := &models.Pedido{}
	id := c.Param("id")
	result := models.Db.Where("pedidos.id = ?", id).Joins("Cliente").Joins("DetallePedidos").First(pedido)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Pedido no encontrado"), nil, c, http.StatusNotFound)
			return
		}

		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener pedido"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, pedido, c, http.StatusOK)
}

func DeletePedido(c *gin.Context) {
	id := c.Param("id")
	pedido := &models.Pedido{}
	result := models.Db.Model(pedido).Where("pedidos.id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar pedido"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Omit("contrasena").Where("id = ?", id).First(pedido)
	utils.CrearRespuesta(nil, pedido, c, http.StatusOK)
}
