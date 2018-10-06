package model

import (
	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg"
	"log"
)

type Subscriber struct {
	tableName struct{} `sql:"subscriber"`

	ID string `json:"id" sql:"id, pk"`
	Email string `json:"email" sql:"email"`
	Name string `json:"name"`
	// Email of subscribedTo Person
	SubscriptionMail string `json:"subscription_mail" sql:"subscription_mail"`
}

func CreateSubscriberTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createTableError := db.CreateTable(&Subscriber{}, options)
	if createTableError != nil {
		log.Printf("Error in Creating Table")
		return createTableError
	}
	log.Printf("Table Created Successfully")
	return nil
}

func (subItem* Subscriber) Add(db *pg.DB) error {
	insertError := db.Insert(subItem)
	if insertError != nil {
		return insertError
	}
	return nil
}

func GetAll(db *pg.DB) ([]Subscriber, error) {
	var subs []Subscriber
	_, err := db.Query(&subs, `SELECT * FROM subscriber`)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func Get(db *pg.DB, email string) (Subscriber, error) {
	var sub Subscriber
	_, fetchError := db.QueryOne(&sub, `SELECT * FROM subscriber WHERE email = ?`, email)
	if fetchError != nil {
		return Subscriber{}, fetchError
	}
	return sub, nil
}
