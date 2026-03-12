package app

import (
	"sentinel-ai/internal/camera"
	"sentinel-ai/internal/config"
	"sentinel-ai/internal/database"
	"sentinel-ai/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct{
	Config *config.Config
	DB *sqlx.DB
	Router *gin.Engine
}

func New(cfg *config.Config)(*App,error){

	db,err:=database.NewDB(cfg)
	if err!=nil{
		return nil,err
	}

	cameraRepo:=camera.NewRepository(db)
	cameraService:=camera.NewService(cameraRepo)
	cameraHandler:=camera.NewHandler(cameraService)

	r:=router.SetupRouter(cameraHandler)

	app:=&App{
		Config:cfg,
		DB:db,
		Router:r,
	}

	return app,nil
}