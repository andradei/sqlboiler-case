package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/andradei/sqlboiler-case/models"
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

	err = aquaman.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(fmt.Sprintln("no inserting person:", err))
	}

	trident := &models.Thing{
		Name: "Trident of Atlantis",
	}

	err = aquaman.AddThings(context.Background(), db, true, trident)
	if err != nil {
		panic(fmt.Sprintln("no adding:", err))
	}

}
