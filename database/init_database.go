package database

import (
	"github.com/go-pg/pg"
	"log"
)
var DB *pg.DB

func Init() {
	DB = pg.Connect(&pg.Options{
		User:     "yashishdua",
		Password: "yashish",
		Addr: "localhost:5432",
		Database: "gosub",
	})

	if DB == nil {
		log.Printf("DB Not Connected")
	}
	log.Printf("DB Connected Successfully")
}
