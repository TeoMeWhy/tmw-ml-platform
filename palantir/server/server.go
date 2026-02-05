package server

import (
	"palantir/configs"
	"palantir/controllers"

	"github.com/gofiber/fiber/v3"
)

type AppServer struct {
	Ctlr *controllers.PredictController
	App  *fiber.App
}

func NewAppServer(cfg *configs.Config) (*AppServer, error) {

	ctlr, err := controllers.NewPredictController(cfg)

	if err != nil {
		return nil, err
	}

	app := &AppServer{
		Ctlr: ctlr,
		App:  fiber.New(),
	}

	app.SetupRoutes()

	return app, nil
}

func (s *AppServer) SetupRoutes() {
	s.App.Post("/predict", s.Ctlr.PostPrediction)
}
