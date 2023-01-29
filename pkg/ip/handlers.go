package ip

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

type Quiestions struct {
	Id       int
	Question string
}

type Clientusers struct {
	Id   int
	Name string
}

type Answers struct {
	Id          int
	Answer1     string
	Answer2     string
	Answer3     string
	Answer4     string
	QuiestionId int
}

type ViewData struct {
	User     string
	Id       int
	Question string
	Answer1  string
	Answer2  string
	Answer3  string
	Answer4  string
}

type FormData struct {
	Question string
	Answer   string
	Name     string
	idName   string
}

// Home Обработчик главной страницы.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		//http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

// NamePage Обработчик отображение страницы с формой ввода имени.
func NamePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/name" {
		http.NotFound(w, r)
		return
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/name.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		//http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// NextTest Обработчик отображение страницы с формой
func NextTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/next_test" {
		http.NotFound(w, r)
		return
	}

	var dsn = "host=localhost user=postgres password=rootroot dbname=Dandelions port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Нет подключения к БД", err)
	}

	var conclusion []ViewData

	user := Clientusers{
		Name: r.FormValue("name"),
	}
	fmt.Println(user.Name)

	// Создать запись Clientusers
	db.Create(&Clientusers{Name: user.Name})

	// Получить последнею запись
	db.Last(&user)
	fmt.Printf("\nПоследнею запись Clientusers %v  %v\n", user.Name, user.Id)

	data := ViewData{
		User: user.Name,
		Id:   user.Id,
	}
	fmt.Println(data)

	conclusion = append(conclusion, data)
	fmt.Println(conclusion)

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/next_test.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, conclusion)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

// FormTest Обработчик сохранения данных страницы с формой
func FormTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.NotFound(w, r)
		return
	}

	var dsn = "host=localhost user=postgres password=rootroot dbname=Dandelions port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Нет подключения к БД", err)
	}

	// Извлечение всех объектов
	var all []Quiestions
	db.Find(&all)

	// Рандомно выбираем первичный ключ
	strId := randomId(all)

	// Извлечение объекта с помощью первичного ключа
	var question Quiestions
	db.First(&question, strId)
	fmt.Println(question)

	// Извлечение объекта с помощью первичного ключа
	var answer Answers
	db.Where("quiestionid = ?", strId).First(&answer)

	fmt.Println(answer.Answer1, answer.Answer2, answer.Answer3, answer.Answer4)

	fmt.Println("/-------------------------------------------- Прилет")

	user := Clientusers{
		Name: r.FormValue("name"),
	}
	fmt.Println(user.Name)

	form := FormData{
		idName:   r.FormValue("name"),
		Question: r.FormValue("question"),
		Answer:   r.FormValue("answer"),
	}
	fmt.Println(form.Name)
	fmt.Println(form.idName)
	fmt.Println(form.Answer)

	fmt.Println("/---------------------------------------------")

	data := ViewData{
		User:     form.idName,
		Question: question.Question,
		Answer1:  answer.Answer1,
		Answer2:  answer.Answer2,
		Answer3:  answer.Answer3,
		Answer4:  answer.Answer4,
	}
	fmt.Println(data)

	var conclusion []ViewData

	conclusion = append(conclusion, data)

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, conclusion)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}
}

func randomId(all []Quiestions) int {
	sl := make([]int, 0, 60)
	rand.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] }) // Рандом Id

	for _, s := range all {
		sl = append(sl, s.Id)
	}
	fmt.Print("\n", sl)

	return sl[0]
}
