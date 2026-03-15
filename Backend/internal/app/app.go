package app

import (
	"context"
	"sentinel-ai/internal/camera"
	"sentinel-ai/internal/organization"
	"sentinel-ai/internal/config"
	"sentinel-ai/internal/database"
	"sentinel-ai/internal/router"
	"sentinel-ai/internal/worker"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Config *config.Config
	DB     *sqlx.DB
	Router *gin.Engine
	CameraRepo *camera.Repository
}

func New(cfg *config.Config) (*App, error) {

	db, err := database.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	cameraRepo := camera.NewRepository(db)
	cameraService := camera.NewService(cameraRepo)
	cameraHandler := camera.NewHandler(cameraService)

	orgRepo := organization.NewRepository(db)
	orgService := organization.NewService(orgRepo)
	orgHandler := organization.NewHandler(orgService)

	r := router.SetupRouter(cameraHandler, orgHandler)

	app := &App{
		Config: cfg,
		DB:     db,
		Router: r,
		CameraRepo: cameraRepo,
	}

	connWorker := worker.NewConnectivityWorker(cameraRepo)
	go connWorker.Start(context.Background())

	return app, nil
}