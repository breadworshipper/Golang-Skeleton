package cmd

import (
	"mm-pddikti-cms/internal/adapter"
	user "mm-pddikti-cms/internal/module/user/entity"

	"github.com/rs/zerolog/log"
)

func RunMigrate() error {
	adapter.Adapters.Sync(
		adapter.Postgres(),
	)

	err := adapter.Adapters.Postgres.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}

	log.Info().Msg("Migration Success")
	return nil
}
