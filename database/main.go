package database

import (
	"github.com/go-pg/pg"
	SubscriberModel "github.com/gosubscriber/model"
	"log"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "yashishdua",
		Password: "yashish",
		Addr: "localhost:5432",
		Database: "gosub",
	})

	if db == nil {
		log.Printf("DB Not Connected")
	}
	log.Printf("DB Connected Successfully")

	SubscriberModel.CreateSubscriberTable(db)

	return db
}