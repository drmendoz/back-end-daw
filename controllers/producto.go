package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductos(c *gin.Context) {
	productos := []*models.Producto{}
	err := models.Db.Where("estado = ?", true).Find(&productos).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener productos"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, productos, c, http.StatusOK)
}

func CreateProducto(c *gin.Context) {
	producto := &models.Producto{}
	err := c.ShouldBindJSON(producto)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	result := models.Db.Create(producto)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear producto"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, producto, c, http.StatusCreated)

}

func GetProductoPorId(c *gin.Context) {
	producto := &models.Producto{}
	id := c.Param("id")
	result := models.Db.Where("id = ? and estado= ?", id, true).First(producto)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Produco no encontrado"), nil, c, http.StatusNotFound)
			return
		}
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener producto"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, producto, c, http.StatusOK)
}

func DeleteProdcuto(c *gin.Context) {
	id := c.Param("id")
	producto := &models.Producto{}
	result := models.Db.Model(producto).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar producto"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Where("id = ?", id).First(producto)
	utils.CrearRespuesta(nil, producto, c, http.StatusOK)
}

func UpdateProducto(c *gin.Context) {

	producto := &models.Producto{}

	err := c.ShouldBindJSON(producto)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	result := models.Db.Where("id = ?", id).Updates(producto)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar prodcuto"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, producto, c, http.StatusOK)

}
