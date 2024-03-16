package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tacoProject/internal/services"
)

type TacoController struct {
	tacoService services.TacoService
}

func NewTacoController(service services.TacoService) *TacoController {
	return &TacoController{
		tacoService: service,
	}
}

func (controller *TacoController) GetById(ctx *gin.Context) {
	product, _ := controller.tacoService.FindById(ctx.GetInt64("id"))
	ctx.JSON(http.StatusOK, product)
}

func (controller *TacoController) GetByName(ctx *gin.Context) {
	taco, _ := controller.tacoService.FindByName(ctx.Param("name"))
	ctx.JSON(http.StatusOK, taco)
}

func (controller *TacoController) CreateTacoByProducts(ctx *gin.Context) {
	taco, err := controller.tacoService.CreateTacoByProducts()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, taco)

}
