package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	ip "oduvanchiki/pkg/ip"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	PORT := ":4000"

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./static".
	var fs = http.FileServer(http.Dir("./static/"))

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	srv := &http.Server{
		Handler:      router,
		Addr:         PORT,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	router.HandleFunc("/", ip.Home).Methods("GET")
	router.HandleFunc("/name", ip.NamePage).Methods("GET")
	router.HandleFunc("/next_test", ip.NextTest).Methods("GET")
	//router.HandleFunc("/", ip.DisplayData).Methods("GET")

	log.Print("Запуск сервера на http://127.0.0.1", PORT)

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	graceShutdown(srv)

}

// Выключает сервер
func graceShutdown(srv *http.Server) {
	interaptCh := make(chan os.Signal, 1)
	signal.Notify(interaptCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interaptCh
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		return
	}
	log.Printf("Выключение")
	os.Exit(0)
}
