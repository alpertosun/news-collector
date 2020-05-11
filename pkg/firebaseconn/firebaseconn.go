package firebaseconn

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

func CreateNewConnection() (*db.Client,context.Context) {
	var opt = option.WithCredentialsFile("credentials.json")
	var ctx = context.Background()
	var conf = &firebase.Config{DatabaseURL: "https://somefirebaseapp.firebaseio.com/"}
	var app, err = firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatal("Error connecting app:",err)
	}
	firebaseDB, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}
	return firebaseDB,ctx
}