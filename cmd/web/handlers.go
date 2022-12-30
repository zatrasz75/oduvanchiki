package main

import (
	"html/template"
	"log"
	"net/http"
)

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}
func formPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./ui/html/form.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}
