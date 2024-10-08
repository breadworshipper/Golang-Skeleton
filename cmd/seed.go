package cmd

import (
	"flag"
	"mm-pddikti-cms/db/seeds"
	"mm-pddikti-cms/internal/adapter"
	"strings"

	"github.com/rs/zerolog/log"
)

func RunSeed(cmd *flag.FlagSet, args []string) {
	total := cmd.Int("total", 1, "total number of records to seed")

	if err := cmd.Parse(args); err != nil {
		log.Fatal().Err(err).Msg("Error while parsing flags")
	}

	tables := args
	log.Info().Msg(strings.Join(tables, ", "))

	if len(tables) == 0 {
		log.Fatal().Msg("No tables provided for seeding")
	}

	adapter.Adapters.Sync(adapter.Postgres())
	defer func() {
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
	}()

	log.Info().Msgf("Seeding tables %s with %d records each", strings.Join(tables, ", "), *total)
	for _, table := range tables {
		log.Info().Msgf("Seeding table %s", table)
		seeds.Execute(adapter.Adapters.Postgres, table, *total)
	}
}
