package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCategorias(c *gin.Context) {
	categorias := []*models.Categoria{}
	err := models.Db.Where("estado = ?", true).Find(&categorias).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener categorias"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, categorias, c, http.StatusOK)
}

func CreateCategoria(c *gin.Context) {
	cat := &models.Categoria{}
	err := c.ShouldBindJSON(cat)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	result := models.Db.Create(cat)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear categoria"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, cat, c, http.StatusCreated)

}

func GetCategoriaPorId(c *gin.Context) {
	cat := &models.Categoria{}
	id := c.Param("id")
	result := models.Db.Where("id = ? and estado= ?", id, true).First(cat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Categoria no encontrada"), nil, c, http.StatusNotFound)
			return
		}
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener categoria"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, cat, c, http.StatusOK)
}

func DeleteCategoria(c *gin.Context) {
	id := c.Param("id")
	cat := &models.Categoria{}
	result := models.Db.Model(cat).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Where("id = ?", id).First(cat)
	utils.CrearRespuesta(nil, cat, c, http.StatusOK)
}

func UpdateCategoria(c *gin.Context) {

	cat := &models.Categoria{}

	err := c.ShouldBindJSON(cat)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	result := models.Db.Where("id = ?", id).Updates(cat)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, cat, c, http.StatusOK)

}
