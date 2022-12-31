package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"oduvanchiki/pkg"
	"os"
)

// Инициализируем FileServer, он будет обрабатывать
// HTTP-запросы к статическим файлам из папки "./ui/static".
var (
	fs = http.FileServer(http.Dir("./ui/static/"))
)

func main() {

	PORT := ":4000"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Использование номера порта по умолчанию: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}
	router := mux.NewRouter()
	//	mux := http.NewServeMux()

	router.HandleFunc("/", pkg.Home).Methods("GET")
	router.HandleFunc("/form", pkg.FormPage).Methods("GET")
	router.HandleFunc("/form", pkg.FormSave).Methods("POST")

	// Обработка всех url будет происходить через router
	http.Handle("/", router)

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/".
	router.Handle("/static/", http.StripPrefix("/static", fs))

	log.Print("Запуск сервера на http://127.0.0.1", PORT)
	err := http.ListenAndServe(PORT, router)
	log.Fatal(err)
}
