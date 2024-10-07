package adapter

import (
	// "log"

	"fmt"
	"pddikti-cms/internal/infrastructure/config"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgres struct {
	adapter *Adapter
}

func Postgres() Option {
	return &postgres{}
}

func (p *postgres) Start(a *Adapter) {
	dbUser := config.Envs.Postgres.Username
	dbPassword := config.Envs.Postgres.Password
	dbName := config.Envs.Postgres.Database
	dbHost := config.Envs.Postgres.Host
	dbSSLMode := config.Envs.Postgres.SslMode
	dbPort := config.Envs.Postgres.Port

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)
	db, err := gorm.Open(pg.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to Postgres")
	}

	// check connection
	err = db.Exec("SELECT 1").Error
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to PDDikti CMS Postgres")
	}

	a.Postgres = db
	p.adapter = a
	log.Info().Msg("PDDikti CMS Postgres connected")
}

func (p *postgres) Close() error {
	db, err := p.adapter.Postgres.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}
	log.Info().Msg("PDDikti CMS Postgres disconnected")

	return nil
}
