package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/shinofara/alpha/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)

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
