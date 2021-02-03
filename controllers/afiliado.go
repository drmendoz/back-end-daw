package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAfiliados(c *gin.Context) {
	afiliados := []*models.Afiliado{}
	err := models.Db.Omit("clave").Where("afiliados.estado = ?", true).Joins("Categoria").Find(&afiliados).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener afiliados"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, afiliados, c, http.StatusOK)
}

func CreateAfiliados(c *gin.Context) {
	afi := &models.InsertAfiliado{}
	err := c.ShouldBindJSON(afi)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	result := models.Db.Create(afi)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear afiliado"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, afi, c, http.StatusCreated)

}

func GetAfiliadoPorId(c *gin.Context) {
	afi := &models.Afiliado{}
	id := c.Param("id")
	result := models.Db.Omit("clave").Where("afiliados.id = ? and afiliados.estado= ?", id, true).Joins("Categoria").First(afi)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Afiliado no encontrado"), nil, c, http.StatusNotFound)
			return
		}
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener afiliado"), nil, c, http.StatusInternalServerError)
		return
	}

	utils.CrearRespuesta(nil, afi, c, http.StatusOK)
}

func DeleteAfiliado(c *gin.Context) {
	id := c.Param("id")
	afi := &models.Afiliado{}
	result := models.Db.Model(afi).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar afiliado"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Where("id = ?", id).First(afi)
	utils.CrearRespuesta(nil, afi, c, http.StatusOK)
}

func UpdateAfiliado(c *gin.Context) {

	afi := &models.InsertAfiliado{}

	err := c.ShouldBindJSON(afi)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	result := models.Db.Where("id = ?", id).Updates(afi)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar afiliado"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, afi, c, http.StatusOK)

}
