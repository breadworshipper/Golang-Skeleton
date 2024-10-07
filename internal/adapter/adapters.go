package adapter

import (
	"fmt"
	"net/http"
	"strings"

	// import "monster-laut-depok/internal/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	Adapters 	*Adapter
	options    	[]Option
)

type Option interface {
	Start(a *Adapter)
	Close() error
}

type Validator interface {
	Validate(i any) []ErrorResponse
}

type Adapter struct {
	// Driving Adapters
	RestServer *fiber.App
	WsServer   *http.Server

	//Driven Adapters
	Postgres  *gorm.DB
	Validator Validator
	Storage   *minio.Client
}

func (a *Adapter) Sync(opts ...Option) {
	for _, o := range opts {
		o.Start(a)
		options = append(options, o)
	}
}

func (a *Adapter) Unsync() error {
	var errs []string

	for _, o := range options {
		if err := o.Close(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		err := fmt.Errorf(strings.Join(errs, "\n"))
		log.Error().Msgf("Error while disconnecting adapters: %v", err)
		return err
	}

	return nil
}
