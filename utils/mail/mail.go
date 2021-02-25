package mail

import (
	"net/smtp"

	"github.com/drmendoz/backend-gin-gorm/utils"
)

var usuario string
var contrasena string

type Contacto struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Detalles string `json:"detalles"`
}

var auth smtp.Auth

func init() {
	usuario = utils.Viper.GetString("MAIL_USER")
	contrasena = utils.Viper.GetString("MAIL_PASSWORD")
	auth = smtp.PlainAuth("", usuario, contrasena, "smtp.gmail.com")
}

func EnviarCorreoContacto(contacto Contacto) error {
	r := NewRequest([]string{"juanjoloor@hotmail.com"}, "WaffleStop", "Nuevo contacto!")
	err := r.ParseTemplate("views/mail.html", contacto)
	if err != nil {
		return err
	}
	_, err = r.SendEmail()
	if err != nil {
		return err
	}
	return err
}
