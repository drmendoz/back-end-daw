package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/drmendoz/backend-gin-gorm/auth"
	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
)

func GetAdministradores(c *gin.Context) {
	administradores := []*models.Administrador{}
	err := models.Db.Omit("usuarios.contrasena").Where("administradores.estado = ?", true).Joins("Usuario").Find(&administradores).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener administadores"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, administradores, c, http.StatusOK)
}

func CreateAdministrador(c *gin.Context) {
	adm := &models.Administrador{}
	err := c.ShouldBindJSON(adm)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	clave := auth.HashPassword(adm.Usuario.Contrasena)
	adm.Usuario.Contrasena = clave
	result := models.Db.Create(adm)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, adm, c, http.StatusCreated)

}

func UpdateAdministrador(c *gin.Context) {

	adm := &models.Administrador{}

	err := c.ShouldBindJSON(adm)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	ui, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	adm.ID = uint(ui)
	adm.Usuario.Contrasena = auth.HashPassword(adm.Usuario.Contrasena)
	result := models.Db.Updates(adm)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, adm, c, http.StatusOK)

}

func GetAdministradorPorId(c *gin.Context) {
	adm := &models.Administrador{}
	id := c.Param("id")
	result := models.Db.Where("administradores.id = ?", id).Omit("usuarios.contrasena").Joins("Usuario").First(adm)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Administrador no encontrado"), nil, c, http.StatusNotFound)
			return
		}

		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, adm, c, http.StatusOK)
}

func DeleteAdministrador(c *gin.Context) {
	id := c.Param("id")
	adm := &models.Administrador{}
	result := models.Db.Model(adm).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Omit("contrasena").Where("id = ?", id).First(adm)
	utils.CrearRespuesta(nil, adm, c, http.StatusOK)
}
