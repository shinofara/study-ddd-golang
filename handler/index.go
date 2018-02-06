package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"gitlab.com/shinofara/alpha/domain/post"
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

	c := post.New(client, ctx)
	c.Set("aaa", &post.Post{
		Text:   "hoge",
		UserID: 1,
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

	fmt.Fprintln(w, "ok")

	user, err := c.Find("aaa")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%+v", user)
}
