package app

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"simpleApi/config"
	"simpleApi/pkg/govalidator"
)

func Run() {
	validator := govalidator.New()

	log := zerolog.New(os.Stdout)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = validator.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log = log.Level(zerolog.Level(*cfg.Logger.Level)).With().Timestamp().Logger()

	defer log.Info().Msg("Application has been shut down")

	log.Debug().Msg("Loaded configuration")

}
