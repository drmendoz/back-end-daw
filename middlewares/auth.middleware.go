package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

//Credentials to login
type Credentials struct {
	Contrasena        string `json:"contrasena"`
	Correo            string `json:"correo"`
	TokenNotificacion string `json:"token_notificacion,omitempty"`
}

//Claims lo que se guarda en los tokens
type Claims struct {
	Rol string
	Id  int
	jwt.StandardClaims
}

func GenerarToken(cliente models.Cliente) string {
	expirationTime := time.Now().Add(200 * time.Minute)
	claims := &Claims{
		Rol: "cliente",
		Id:  int(cliente.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			err := errors.New("No posee token de autorizacion")
			_ = c.Error(err)
			utils.CrearRespuesta(err, nil, c, http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || err == jwt.ErrSignatureInvalid {
			err := c.Error(errors.New("Token Invalido"))
			_ = c.Error(err)
			utils.CrearRespuesta(err, nil, c, http.StatusUnauthorized)
			c.Abort()
			return

		}
		if !tkn.Valid {
			err := c.Error(errors.New("Token Expirado"))
			_ = c.Error(err)
			utils.CrearRespuesta(err, nil, c, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("usuario", claims)

		c.Next()

	}
}
