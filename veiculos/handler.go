package veiculos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{
	storage MySQLStorage
}

func NewVeiculo(stg MySQLStorage) *Controller {
	return &Controller{
		storage: stg
	}
}

//endpoint que busca os veiculos
func (ctrl *Controller) Get(c *gin.Context) {
	veiculos, err := ctrl.storage.GetVeiculos()
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, veiculos)
}

//endpoint que cria novos veiculos
func (ctrl * Controller) Create(c *gin.Context) {
	var v Veiculo
	//transforma a request em um objeto do tipo Veiculo
	if err := c.ShouldBindJSON(&v); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//salva os dados no banco
	err := ctrl.storage.CreateVeiculo(v.Nome, v.Marca, v.Ano, v.Modelo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

//atualiza um veiculo
func (ctrl * Controller) Update(c *gin.Context) {
	var v Veiculo
	//transforma a request em um objeto do tipo Veiculo
	if err := c.ShouldBindJSON(&v); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//salva os dados no banco
	err := ctrl.storage.UpdateVeiculo(v.id, &v)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

//apaga um veiculo
func (ctrl * Controller) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//declara a variavel e ao mesmo tempo verifica se é diferente de nil
	if err := ctrl.storage.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
