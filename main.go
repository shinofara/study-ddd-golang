package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/shinofara/alpha/handler"
	"gitlab.com/shinofara/alpha/infrastructure/locator"
)

func main() {

	l := locator.New()
	l.SetStorage()

	http.HandleFunc("/", handler.Index(l.ServiceLocator()))

	srv := &http.Server{Addr: ":28080"}

	// サーバはブロックするので別の goroutine で実行する
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// シグナルを待つ
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh

	// シグナルを受け取ったらShutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
