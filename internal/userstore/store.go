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
	_, err := config.NewConfig()
	if err != nil {
		log.Println("can't get config", err)
		return nil
	}

	pgurl := "postgres://postgres:Flints1488@ec2-13-51-242-14.eu-north-1.compute.amazonaws.com:5432/users?sslmode=disable"

	conn, err := NewPG(ctx, pgurl)
	if err != nil {
		log.Println("can't initialize Postgres", err)
	}
	if err = conn.db.Ping(ctx); err != nil {
		log.Println("can't connect to DataBase", err)
	}

	return &Store{
		PG:   conn,
		user: model.NewUser(),
	}
}
