package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andradei/sqlboiler-case/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, err := sql.Open("postgres", `postgres://postgres:password@localhost:5432/postgres?sslmode=disable`)
	if err != nil {
		panic(fmt.Sprintln("no opening:", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintln("no pinging:", err))
	}

	aquaman := models.Person{
		Name: "Aquaman",
		Age:  32,
	}

	trident := &models.Thing{
		Name: "Trident of Atlantis",
	}

	err = trident.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(fmt.Sprintln("no inserting thing:", err))
	}
	err = aquaman.AddThings(context.Background(), db, false, trident)
	if err != nil {
		panic(fmt.Sprintln("no adding:", err))
	}

	err = aquaman.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(fmt.Sprintln("no inserting person:", err))
	}
}
