package webserver

import (
	"github.com/gin-gonic/gin"
)

//função externa
func New() *gin.Engine {
	return startServer(gin.New())
}

//função interna
func startServer(r *gin.Engine) *gin.Engine {
	//agrupa os endpoints
	v1 := r.Group("api/v1")
	//configurar endpoint do veiculo
	handler := veiculo.NewVeiculo(stgVeiculo)
	v1.GET("/veiculos", handler.Get)
	v1.POST("/veiculos", handler.Create)
	v1.PUT("/veiculos", handler.Update)
	//passagem de parametro url = http://localhost:8080/api/v1/veiculos/1 => ID
	v1.DELETE("/veiculos/:id", handler.Delete)

	return r
}

func CreateDB() veiculo.MySQLStorage {
	conn := "root:root@tcp(127.0.0.1)/veiculos"
	return veiculo.NewStorage(conn)
}
