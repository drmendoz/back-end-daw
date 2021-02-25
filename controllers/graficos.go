package controllers

import (
	"errors"
	"net/http"

	"github.com/drmendoz/backend-gin-gorm/models"
	"github.com/drmendoz/backend-gin-gorm/utils"
	"github.com/gin-gonic/gin"
)

type Diagramas struct {
	Ciudades   []string `json:"ciudades"`
	Categorias []int64  `json:"categorias"`
}

func GetCiudadesGrafico(c *gin.Context) {
	var results []map[string]interface{}

	err := models.Db.Raw(" select ciudad, count(*) as cantidad FROM pedidos GROUP BY ciudad;").Scan(&results).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener pedidos"), nil, c, http.StatusInternalServerError)
		return
	}

	utils.CrearRespuesta(err, results, c, http.StatusOK)

}

func GetProductosGrafico(c *gin.Context) {

	var results []map[string]interface{}
	err := models.Db.Raw("  select distinct(productos.nombre), count(*) as c FROM detalle_pedidos join productos on productos.id=detalle_pedidos.producto_id GROUP BY productos.nombre;").Scan(&results).Error
	if err != nil {
		_ = c.Error(err)
		utils.CrearRespuesta(errors.New("Error al obtener productos"), nil, c, http.StatusOK)
		return
	}
	utils.CrearRespuesta(err, results, c, http.StatusOK)
}

func GetClientesGrafico(c *gin.Context) {
	var clientes [10]int64
	clientes[0] = 0
	clientes[1] = 4
	clientes[2] = 3
	clientes[3] = 4
	clientes[5] = 17
	clientes[6] = 10
	clientes[7] = 30
	clientes[8] = 44
	clientes[9] = 33
	utils.CrearRespuesta(nil, clientes, c, http.StatusOK)
}
