package userstore

import (
	"context"
	"log"
	"websirk/config"
	"websirk/internal/model"
)

type Store struct {
	PG   *postgres
	user *model.User
}

func NewStore() *Store {

	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("can't get config", err)
		return nil
	}

	conn, err := NewPG(ctx, cfg.URL)
	if err != nil {
		log.Println("can't initialize Postgres", err)
	}
	if err = conn.db.Ping(ctx); err != nil {
		log.Println("can't connect to c2c DataBase", err)
	}

	return &Store{
		PG:   conn,
		user: model.NewUser(),
	}
}
