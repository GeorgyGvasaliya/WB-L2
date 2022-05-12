package main

import (
	"L2/develop/11/internal/application"
	"L2/develop/11/internal/config"
	"L2/develop/11/internal/server/httpserver"
	"L2/develop/11/internal/storage/memory"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	storage := memory.New()
	app := application.NewCalendar(storage)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := httpserver.New(cfg.HTTP, app)
	go func() {
		log.Println("http server start...")
		if err := server.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("http server http stopped....")
			} else {
				log.Fatalln(err)
			}
		}
	}()
	<-stop
	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Shutdown(ctxClose)
	if err != nil {
		log.Fatalln(err)
	}
}
