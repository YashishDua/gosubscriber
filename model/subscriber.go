package model

import (
	"github.com/go-pg/pg/orm"
	"gosubscriber/database"
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

func CreateSubscriberTable() error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createTableError := database.DB.CreateTable(&Subscriber{}, options)
	if createTableError != nil {
		log.Printf("Error in Creating Table")
		return createTableError
	}
	log.Printf("Table Created Successfully")
	return nil
}

func (subItem* Subscriber) Add() error {
	insertError := database.DB.Insert(subItem)
	if insertError != nil {
		return insertError
	}
	return nil
}

func GetAll() ([]Subscriber, error) {
	var subs []Subscriber
	_, err := database.DB.Query(&subs, `SELECT * FROM subscriber`)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func Get(email string) (Subscriber, error) {
	var sub Subscriber
	_, fetchError := database.DB.QueryOne(&sub, `SELECT * FROM subscriber WHERE email = ?`, email)
	if fetchError != nil {
		return Subscriber{}, fetchError
	}
	return sub, nil
}
