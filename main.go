package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"oduvanchiki/pkg/ip"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var errlog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var inflog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

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
	router.HandleFunc("/next_test", ip.NextTest).Methods("POST")
	router.HandleFunc("/test", ip.FormTest).Methods("POST")

	inflog.Print("Запуск сервера на http://127.0.0.1", PORT)

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			inflog.Println(err)
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
		inflog.Printf("Ошибка при закрытии прослушивателей или тайм-аут контекста %v", err)
		return
	}
	inflog.Printf("Выключение сервера")
	os.Exit(0)
}
