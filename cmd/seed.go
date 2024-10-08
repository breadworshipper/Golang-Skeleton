package cmd

import (
	"flag"
	"mm-pddikti-cms/db/seeds"
	"mm-pddikti-cms/internal/adapter"

	"github.com/rs/zerolog/log"
)

func RunSeed(cmd *flag.FlagSet, args []string) {
	var (
		total = cmd.Int("total", 1, "total of records to seed")
	)

	if err := cmd.Parse(args); err != nil {
		log.Fatal().Err(err).Msg("Error while parsing flags")
	}

	adapter.Adapters.Sync(
		adapter.Postgres(),
	)
	defer func() {
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
	}()

	var table *string
	for _, v := range args {
		table = cmd.String("table", v, "seed to run")
		seeds.Execute(adapter.Adapters.Postgres, *table, *total)
	}
}
