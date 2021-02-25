package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/drmendoz/backend-gin-gorm/utils/mail"
	"github.com/gin-gonic/gin"
)

func HandleContactenos(c *gin.Context) {
	contacto := &mail.Contacto{}
	err := c.ShouldBindJSON(contacto)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	err = mail.EnviarCorreoContacto(*contacto)
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al enviar solicitud"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, "Correo enviado satisfactoriamente", c, http.StatusOK)

}
