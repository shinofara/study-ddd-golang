package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"gitlab.com/shinofara/alpha/domain/data/message"
	"gitlab.com/shinofara/alpha/domain/data/user"
	serviceCh "gitlab.com/shinofara/alpha/domain/service/channel"
	serviceMess "gitlab.com/shinofara/alpha/domain/service/message"
	infraCh "gitlab.com/shinofara/alpha/infrastructure/firestore/channel"
	infraMess "gitlab.com/shinofara/alpha/infrastructure/firestore/message"
	infraUser "gitlab.com/shinofara/alpha/infrastructure/firestore/user"
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

	// action内で使用するrepositoryを初期化
	userRepo := infraUser.New(client, ctx)
	messRepo := infraMess.New(client, ctx)
	channelRepo := infraCh.New(client, ctx)

	// owner作成
	userService := user.NewService(userRepo)
	u, err := userService.Register("しのはら")
	if err != nil {
		panic(err)
	}

	// channel新規作成
	chService := serviceCh.New(channelRepo, userRepo, messRepo)
	ch, err := chService.Create("テスト", u)
	if err != nil {
		panic(err)
	}

	// channelに投稿
	messService := serviceMess.New(messRepo)
	messSpec := &message.PostSpecification{MinLength: 1, MaxLength: 100}

	mess, err := messService.Post(ch.ID, u.ID, "初投稿", messSpec)
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
