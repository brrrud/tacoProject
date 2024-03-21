package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/controllers"
	"tacoProject/internal/repositories"
	"tacoProject/internal/services"
)

func Initialize(router *gin.Engine, db *pgx.Conn) {
	//Repository
	tacoRepository := repositories.NewTacoRepository(db)
	pfcInfoRepository := repositories.NewPfcInfoRepository(db)

	//Service
	tacoService := services.NewTacoServiceImpl(tacoRepository)
	pfcInfoService := services.NewPfcInfoServiceImpl(pfcInfoRepository)

	//Controller
	tacoController := controllers.NewTacoController(tacoService)
	pfcInfoController := controllers.NewPfcInfoController(pfcInfoService)

	router.GET("/byId/:id", tacoController.GetById)
	router.GET("/byName/:name", tacoController.GetByName)
	router.POST("/createByProducts", tacoController.CreateTacoByProducts)
	router.POST("/createPFC", pfcInfoController.PostPfcInfo)
}
