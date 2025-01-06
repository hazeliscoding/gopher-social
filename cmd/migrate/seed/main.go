package main

import (
	"log"

	"github.com/hazeliscoding/gopher-social/internal/db"
	"github.com/hazeliscoding/gopher-social/internal/env"
	"github.com/hazeliscoding/gopher-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://postgres:postgres@localhost:5432/gopher_social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
