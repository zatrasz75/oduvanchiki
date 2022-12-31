package main

import (
	"html/template"
	"log"
	"net/http"
)

// Используем функцию template.ParseFiles() для чтения файлов шаблона.
var (
	tmpl = template.Must(template.ParseFiles("./ui/html/form.html"))
)

// Обработчик главной страницы.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к трем файлам. Обратите внимание, что
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

// Обработчик отображение страницы с формой
func formPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	data := ViewData{
		Title: "Одуванчики",
		Users: []string{"Krex", "Pex", "Fex"},
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// Обработчик сохранения данных страницы с формой
func formSave(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// https://www.youtube.com/watch?v=lKvQYHZtuzA&list=PLHUicSITKZEmz2w3zo-aUpxCUZuqONE4c&index=2
