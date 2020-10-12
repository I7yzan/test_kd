package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//"gitlab.com/conlist/pgdb"
	"log"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://contacts:ooshi9si@172.16.122.22:15432/contacts?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//err = pgdb.CreateOwner(db, 1001, 101, 101, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
