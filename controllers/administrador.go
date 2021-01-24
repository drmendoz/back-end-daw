package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/drmendoz/backend-gin-gorm/auth"
	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
)

func GetAdministradores(c *gin.Context) {
	administradores := []*models.Administrador{}
	err := models.Db.Find(&administradores).Error
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
	clave := auth.HashPassword(adm.Contrasena)
	adm.Contrasena = clave
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
	adm.Contrasena = auth.HashPassword(adm.Contrasena)
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
	result := models.Db.Where("id = ?", id).Omit("contrasena").First(adm)
	if result.Error != nil {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener administrador"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(nil, adm, c, http.StatusOK)
}
