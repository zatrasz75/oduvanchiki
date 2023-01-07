package ip

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oduvanchiki/pkg/db"
)

type ViewData struct {
	Title    string
	Answers  []string
	Question string
	Answer1  string
	Answer2  string
	Answer3  string
	Answer4  string
}

type FormData struct {
	Id          string
	Question    string
	AnswerTrue  string
	AnswerFals1 string
	AnswerFals2 string
	AnswerFals3 string
}

// Используем функцию template.ParseFiles() для чтения файлов шаблона.
var (
	tmpl = template.Must(template.ParseFiles("./static/html/form.html"))
)

// Home Обработчик главной страницы.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к трем файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./static/html/home.page.tmpl",
		"./static/html/base.layout.tmpl",
		"./static/html/footer.partial.tmpl",
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

// FormPage Обработчик отображение страницы с формой
func FormPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	var strId = 1
	slId := make([]int, 0, 60)

	quiest := db.QuiestDB()
	for _, s := range quiest {
		fmt.Println(s.Id, s.Question)
		slId = append(slId, s.Id)
	}
	fmt.Println(slId) // Массив из ID нид сделать из него рандом и записать в strId

	answ := db.AnswersDB()
	for _, s := range answ {
		fmt.Println(s.Id, s.Answer1, s.Answer2, s.Answer3, s.Answer4)
	}

	q := db.QuiestionOneDB(strId)
	a := db.AnswerOneDB(strId)

	data := ViewData{
		Title:    "Одуванчики",
		Question: q.Question,
		Answer1:  a.Answer1,
		Answer2:  a.Answer2,
		Answer3:  a.Answer3,
		Answer4:  a.Answer4,
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

// FormSave Обработчик сохранения данных страницы с формой
func FormSave(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	f := FormData{
		Id:          r.FormValue("Id"),
		AnswerTrue:  r.FormValue("AnswerTrue"),
		AnswerFals1: r.FormValue("AnswerFals1"),
		AnswerFals2: r.FormValue("AnswerFals2"),
		AnswerFals3: r.FormValue("AnswerFals3"),
	}

	data := ViewData{
		Title:   "Одуванчики",
		Answers: []string{"Krex", "Pex", "Fex", "upssssss", "FormSave"},
	}
	if f.AnswerTrue == "Krex" {
		fmt.Println(f)
	}
	//	fmt.Println(f)

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// DisplayData Обработчик отображение страницы с формой
func DisplayData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	data := ViewData{
		Title:   "Одуванчики",
		Answers: []string{"Krex", "Pex", "Fex", "DisplayData"},
	}
	a := db.QuiestionOneDB(1)
	disp := db.Quiestions{
		Id:       a.Id,
		Question: a.Question,
	}
	fmt.Println(disp)
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// https://www.youtube.com/watch?v=lKvQYHZtuzA&list=PLHUicSITKZEmz2w3zo-aUpxCUZuqONE4c&index=2
