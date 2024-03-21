package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tacoProject/internal/models"
	"tacoProject/internal/services"
)

type PfcInfoController struct {
	pfcInfoService services.PfcInfoService
}

func NewPfcInfoController(pfcInfoService services.PfcInfoService) *PfcInfoController {
	return &PfcInfoController{pfcInfoService: pfcInfoService}
}

func (controller *PfcInfoController) PostPfcInfo(ctx *gin.Context) {
	var newPfcInfo models.PfcInfo
	if err := ctx.ShouldBindJSON(&newPfcInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.pfcInfoService.CreateProduct(newPfcInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pfcInfo"})
		return
	}
}
