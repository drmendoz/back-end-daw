package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/middlewares"
	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetResenas(c *gin.Context) {
	ciudad := c.Query("ciudad")
	resenas := []*models.Resena{}
	claims, exists := c.Get("usuario")
	if !exists {
		utils.CrearRespuesta(errors.New("Error al identificar al usuario"), nil, c, http.StatusUnauthorized)
		return

	}
	usuario := claims.(*middlewares.Claims)

	err := models.Db.Where("estado = ? and cliente_id = ? and ciudad like '%"+ciudad+"%'", true, usuario.Id).Find(&resenas).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener resenas"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, resenas, c, http.StatusOK)
}
func GetResenasAdmin(c *gin.Context) {
	ciudad := c.Query("ciudad")
	resenas := []*models.Resena{}

	err := models.Db.Where("estado = ?  and ciudad like '%"+ciudad+"%'", true).Preload("Cliente").Preload("Cliente.Usuario").Find(&resenas).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener resenas"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, resenas, c, http.StatusOK)

}

func CreateResena(c *gin.Context) {
	resena := &models.Resena{}
	err := c.ShouldBindJSON(resena)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	claims, exists := c.Get("usuario")
	if !exists {
		utils.CrearRespuesta(errors.New("Error al identificar al usuario"), nil, c, http.StatusUnauthorized)
		return

	}
	usuario := claims.(*middlewares.Claims)
	resena.ClienteID = usuario.Id

	result := models.Db.Create(resena)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear resena"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, resena, c, http.StatusCreated)

}

func GetResenaPorId(c *gin.Context) {
	resena := &models.Resena{}
	id := c.Param("id")
	result := models.Db.Where("resenas.id = ? and resenas.estado= ?", id, true).Preload("Cliente").Preload("Cliente.Usuario").First(resena)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Resena no encontrado"), nil, c, http.StatusNotFound)
			return
		}
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener resena"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, resena, c, http.StatusOK)
}

func DeleteResena(c *gin.Context) {
	id := c.Param("id")
	resena := &models.Resena{}
	result := models.Db.Model(resena).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar resena"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Where("id = ?", id).First(resena)
	utils.CrearRespuesta(nil, resena, c, http.StatusOK)
}

func UpdateResena(c *gin.Context) {

	resena := &models.Resena{}

	err := c.ShouldBindJSON(resena)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	result := models.Db.Where("id = ?", id).Updates(resena)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar resena"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, resena, c, http.StatusOK)

}
