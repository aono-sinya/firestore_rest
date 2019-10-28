package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"firestore_rest/condGacha"
	"firestore_rest/gacha"
	"firestore_rest/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(body[:length], &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatalln(err)
	}
	users := models.Users{
		Ctx: ctx,
		Client: client,
	}
	doc := users.Create(user)
	fmt.Fprintf(w, "%s\n", doc.ID)
}

func User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatalln(err)
	}
	users := models.Users{
		Ctx: ctx,
		Client: client,
	}
	user := users.Find(ps.ByName("userID"))
	fmt.Fprintf(w, "%s\n", user)
}

func main() {
	router := httprouter.New()
	router.POST("/", Index)
	router.GET("/user/:userID", User)

	router.GET("/gacha", gacha.GachaIndex)
	router.GET("/gacha/:userID", gacha.GachaGet)

	router.GET("/cond/gacha/:userID", condGacha.GachaGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
