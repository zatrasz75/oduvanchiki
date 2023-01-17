package ip

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"oduvanchiki/pkg/db"
)

type ViewData struct {
	Title    string
	Id       int
	Question string
	Answer1  string
	Answer2  string
	Answer3  string
	Answer4  string
}

type FormData struct {
	Id       string
	Question string
	Answer1  string
	Answer2  string
	Answer3  string
	Answer4  string
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
	var strId int

	quiest := db.QuiestDB()
	strId = randomId(quiest, strId)

	q := db.QuiestionOneDB(strId)
	a := db.AnswerOneDB(strId)

	data := ViewData{
		Title:    "Одуванчики",
		Id:       q.Id,
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
	var strId int

	quiest := db.QuiestDB()
	strId = randomId(quiest, strId)

	q := db.QuiestionOneDB(strId)
	a := db.AnswerOneDB(strId)

	data := ViewData{
		Title:    "Одуванчики",
		Id:       q.Id,
		Question: q.Question,
		Answer1:  a.Answer1,
		Answer2:  a.Answer2,
		Answer3:  a.Answer3,
		Answer4:  a.Answer4,
	}

	f := FormData{
		Id:      r.FormValue("Id"),
		Answer1: r.FormValue("Answer"),
		Answer2: r.FormValue("Answer"),
		Answer3: r.FormValue("Answer"),
		Answer4: r.FormValue("Answer"),
	}
	fmt.Println(f.Id, f.Answer1, f.Answer2, f.Answer3, f.Answer4)

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
//func DisplayData(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/form" {
//		http.NotFound(w, r)
//		return
//	}
//	data := ViewData{
//		Title:   "Одуванчики",
//		Answers: []string{"Krex", "Pex", "Fex", "DisplayData"},
//	}
//	a := db.QuiestionOneDB(1)
//	disp := db.Quiestions{
//		Id:       a.Id,
//		Question: a.Question,
//	}
//	fmt.Println(disp)
//	// Затем мы используем метод Execute() для записи содержимого
//	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
//	// возможность отправки динамических данных в шаблон.
//	err := tmpl.Execute(w, data)
//	if err != nil {
//		log.Println(err.Error())
//		http.Error(w, "внутренняя ошибка сервера", 500)
//	}
//}

func randomId(quiest []db.Quiestions, strId int) int {
	sl := make([]int, 0, 60)
	rand.Shuffle(len(quiest), func(i, j int) { quiest[i], quiest[j] = quiest[j], quiest[i] }) // Рандом Id

	for _, s := range quiest {
		sl = append(sl, s.Id)
	}
	fmt.Println(sl)

	return strId + sl[0]
}

// https://www.youtube.com/watch?v=lKvQYHZtuzA&list=PLHUicSITKZEmz2w3zo-aUpxCUZuqONE4c&index=2
