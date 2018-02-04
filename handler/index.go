package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"gitlab.com/shinofara/alpha/domain/chat"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Index(w http.ResponseWriter, r *http.Request) {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Fprintf(w, "error initializing app: %v", err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	//自動でIDが振られる
	//_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{

	//任意のIDを触れる
	/*	_, err = client.Collection("users").Doc("aaa").Set(ctx, map[string]interface{}{
			"first":  "cccccc",
			"middle": "Mathison",
			"last":   "Turing",
			"born":   1912,
			"id":     "a",
		})
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}

	*/

	c := chat.New(client, ctx)
	c.Add(&chat.Chat{
		Message: "hoge",
		User:    "username",
	})

	iter := client.Collection("users").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v<br>", err)
		}
		fmt.Fprintf(w, "%v<br>", doc.Data())
	}

	fmt.Fprint(w, "ok")

	user, err := client.Collection("users").Doc("aaa").Get(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%v", user)
}
