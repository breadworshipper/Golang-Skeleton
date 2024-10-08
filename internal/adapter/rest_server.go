package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type restServer struct {
	adapter *Adapter
	app *fiber.App
}

func RestServer(app *fiber.App) Option {
	return &restServer{app:app}
}

func (r *restServer) Start(a *Adapter) {
	a.RestServer = r.app
	r.adapter = a

	log.Info().Msg("Rest server connected")
}

func (r *restServer) Close() error {
	if err := r.adapter.RestServer.Shutdown(); err != nil {
		return err
	}
	log.Info().Msg("Rest server disconnected")

	return nil
}


// func WithRestServer(app *fiber.App) Option {
// 	log.Info().Msg("Rest server connected")
// 	return func(a *Adapter) {
// 		a.RestServer = app
// 	}
// }
