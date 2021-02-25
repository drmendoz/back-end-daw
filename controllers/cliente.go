package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/auth"
	"github.com/drmendoz/backend-gin-gorm/middlewares"
	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCliente(c *gin.Context) {
	clientes := []*models.Cliente{}
	err := models.Db.Omit("usuarios.contrasena").Where("clientes.estado = ?", true).Joins("Usuario").Find(&clientes).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener clientes"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, clientes, c, http.StatusOK)
}

func CreateCliente(c *gin.Context) {
	cliente := &models.Cliente{}
	err := c.ShouldBindJSON(cliente)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	cliente.Usuario.Contrasena = auth.HashPassword(cliente.Usuario.Contrasena)
	err = models.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(cliente).Error
		if err != nil {
			return err
		}
		return nil
	},
	)
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al crear usuario"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, cliente, c, http.StatusCreated)

}

func GetClientePorId(c *gin.Context) {
	cliente := &models.Cliente{}
	id := c.Param("id")
	result := models.Db.Omit("usuarios.contrasena").Where("clientes.id = ? and clientes.estado= ?", id, true).Joins("Usuario").First(cliente)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("Afiliado no encontrado"), nil, c, http.StatusNotFound)
			return
		}
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al obtener afiliado"), nil, c, http.StatusInternalServerError)
		return
	}

	utils.CrearRespuesta(nil, cliente, c, http.StatusOK)
}

func DeleteCliente(c *gin.Context) {
	id := c.Param("id")
	cliente := &models.Cliente{}
	result := models.Db.Model(cliente).Where("id = ?", id).Update("estado", false)
	if result.Error != nil || result.RowsAffected == 0 {
		_ = c.Error(result.Error)
		utils.CrearRespuesta(errors.New("Error al eliminar cliente"), nil, c, http.StatusInternalServerError)
		return
	}
	_ = models.Db.Where("id = ?", id).First(cliente)
	utils.CrearRespuesta(nil, cliente, c, http.StatusOK)
}

func UpdateCliente(c *gin.Context) {

	cliente := &models.Cliente{}

	err := c.ShouldBindJSON(cliente)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	result := models.Db.Where("id = ?", id).Updates(cliente)
	if result.Error != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al actualizar cliente"), nil, c, http.StatusInternalServerError)
		return
	}
	utils.CrearRespuesta(err, cliente, c, http.StatusOK)

}

func LoginCliente(c *gin.Context) {
	login := &auth.Login{}
	err := c.ShouldBindJSON(login)
	if err != nil {
		utils.CrearRespuesta(err, nil, c, http.StatusBadRequest)
		return
	}
	cliente := &models.Cliente{}
	result := models.Db.Where("correo = ?", login.Usuario).Joins("Usuario").First(cliente)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.CrearRespuesta(errors.New("No existe el correo electronico"), nil, c, http.StatusUnauthorized)
			return
		}
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener credenciales"), nil, c, http.StatusInternalServerError)
		return
	}
	login.Contrasena = auth.HashPassword(login.Contrasena)
	if login.Contrasena != cliente.Usuario.Contrasena {
		utils.CrearRespuesta(nil, "Error al validar credenciales", c, http.StatusUnauthorized)
		return
	}
	clienteLog := models.ClienteLog{}
	clienteLog.Token = middlewares.GenerarToken(*cliente)
	clienteLog.Cliente = *cliente
	cliente.Usuario.Contrasena = ""
	utils.CrearRespuesta(err, clienteLog, c, http.StatusOK)

}
