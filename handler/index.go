package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/shinofara/alpha/domain/service"

	firebase "firebase.google.com/go"
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

	// owner作成
	userService := service.NewUser(client, ctx)
	u, err := userService.Register("しのはら")
	if err != nil {
		panic(err)
	}

	// channel新規作成
	chService := service.NewChannel(client, ctx)
	ch, err := chService.Create("テスト", u)
	if err != nil {
		panic(err)
	}

	// channelに投稿
	messService := service.NewMessage(client, ctx)
	mess, err := messService.Post(ch.ID, u.ID, "初投稿")
	if err != nil {
		panic(err)
	}

	// channel内のメッセージ全件取得
	currentCh, err := chService.InitialDisplay(ch.ID)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "チャンネル作成結果<br>")
	fmt.Fprintf(w, "%+v<br>", ch)

	fmt.Fprint(w, "メッセージ投稿結果<br>")
	fmt.Fprintf(w, "%+v<br>", mess)

	fmt.Fprint(w, "チャンネルの初期表示に必要な情報取得結果<br>")
	fmt.Fprintf(w, "%+v<br>", currentCh)
}
