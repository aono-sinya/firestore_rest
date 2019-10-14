package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"log"
)

type User struct {
	Name string `firestore:"name"`
	Pass string `firestore:"pass"`
}

type Users struct {
	Ctx context.Context
	Client *firestore.Client
}

func (u *Users) collection() *firestore.CollectionRef {
	return u.Client.Collection("users")
}

func (u *Users) Create(user User) *firestore.DocumentRef {
	doc, _, err := u.collection().Add(u.Ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func (u *Users) Find(id string) *User {
	doc, err := u.collection().Doc(id).Get(u.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	user := new(User)
	err = mapToStruct(doc.Data(), &user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}