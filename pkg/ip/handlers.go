package ip

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	Id      int
	Answer1 string
	Answer2 string
	Answer3 string
	Answer4 string
	//QuestionId int
}

type Quizes struct {
	Id      int
	Userid  int
	Started time.Time
}

type Correctanswers struct {
	//Id            int
	//QuiestionId   int
	//Answercorrect string
	Correct bool
}

type Results struct {
	Id         int
	Questionid int
	Answerid   int
	Quizid     int
	Answered   time.Time
	Point      int
}

type ViewData struct {
	User string
	//	Id        int
	Question  string
	Answer1   string
	Answer2   string
	Answer3   string
	Answer4   string
	TestStart int
	TestId    string
}

type FormData struct {
	Question  string
	Answer    string
	Name      string
	TestStart string
	User      string
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

	if user.Name != "" {
		fmt.Printf("\nЗапись %v\n", user.Name)

		// Создать запись Clientusers
		//db.Create(&Clientusers{Name: user.Name})
	} else {
		return
	}

	// Получить последнею запись Clientusers
	db.Last(&user)

	//timeT := startTime()

	//Создать запись Quizes
	//db.Create(&Quizes{Userid: user.Id, Started: timeT})

	var numberTest Quizes
	// Получить последнею запись Quizes
	db.Last(&numberTest)

	data := ViewData{
		User:      user.Name,
		TestStart: numberTest.Id,
	}
	//fmt.Println(data)

	conclusion = append(conclusion, data)
	//	fmt.Println(conclusion)

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

	//--------------------------------------

	form := FormData{
		TestStart: r.FormValue("id"),
		Question:  r.FormValue("question"),
		Answer:    r.FormValue("answer"),
	}

	//-----------------------------------------------------

	// Извлечение всех объектов
	var allq []Quiestions
	db.Find(&allq)

	var resR []Results
	// Извлечение объектов, где поле quizid равно TestStart
	db.Where("quizid = ?", form.TestStart).Find(&resR)

	// Рандомно выбираем первичный ключ
	strId := randomId(allq, resR)
	fmt.Printf("Рандомно выбираем первичный ключ %v\n", strId)

	//-------------------------------------------------------

	// Извлечение объекта с помощью первичного ключа
	var question Quiestions
	db.First(&question, strId)

	// Извлечение объектов, где поле quiestionid равно первичному ключу
	var answer Answers
	db.Where("quiestionid = ?", strId).First(&answer)

	// Извлечение объектов, где поле answercorrect равно form.Answer
	var correct Correctanswers
	db.Where("answercorrect = ?", form.Answer).Find(&correct)
	fmt.Printf("Верный ответ %v\n", correct.Correct)

	// Извлечение объектов, где поле id равно form.TestStart
	var quizes Quizes
	db.Where("id = ?", form.TestStart).Find(&quizes)

	var result Results
	// Правильный ответ
	if correct.Correct == true {
		result.Point = 1
	}

	//timeT := startTime()

	//Создать запись Results
	//db.Create(&Results{Questionid: question.Id, Answerid: answer.Id, Quizid: quizes.Id, Answered: timeT, Point: result.Point})

	fmt.Println("/---------------------------------------------")

	data := ViewData{
		TestId:   form.TestStart,
		Question: question.Question,
		Answer1:  answer.Answer1,
		Answer2:  answer.Answer2,
		Answer3:  answer.Answer3,
		Answer4:  answer.Answer4,
	}

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

// Создает рандомно число
func randomId(allq []Quiestions, resR []Results) int {

	slQ := make([]int, 0, 100)
	// Присвоение значений срезу
	for _, v := range allq {
		slQ = append(slQ, v.Id)
	}

	slR := make([]int, 0, 60)
	// Присвоение значений срезу
	for _, v := range resR {
		slR = append(slR, v.Questionid)
	}
	fmt.Println(slQ)
	fmt.Println(slR)

	var shortest, longest *[]int
	if len(slQ) < len(slR) {
		shortest = &slQ
		longest = &slR
	} else {
		shortest = &slR
		longest = &slQ
	}
	// Самый короткий фрагмент в карту
	var m map[int]bool
	m = make(map[int]bool, len(*shortest))
	for _, s := range *shortest {
		m[s] = false
	}
	// Значения из самого длинного фрагмента, которые не существуют на карте
	var diff []int
	for _, s := range *longest {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
			continue
		}
		m[s] = true
	}
	// Значения с карты, которые не были в самом длинном фрагменте
	for s, ok := range m {
		if ok {
			continue
		}
		diff = append(diff, s)
	}
	fmt.Println(diff)

	sl := make([]int, 0, 100)
	rand.Shuffle(len(diff), func(i, j int) { diff[i], diff[j] = diff[j], diff[i] }) // Рандом Id

	for _, v := range diff {
		sl = append(sl, v)
	}

	return sl[0]
}

func startTime() time.Time {

	tNow := time.Now()
	//Время для Unix Timestamp
	tUnix := tNow.Unix()
	//Временная метка Unix для time.Time
	time.Unix(tUnix, 0)

	return time.Now()
}
