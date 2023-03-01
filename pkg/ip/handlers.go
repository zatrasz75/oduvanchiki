package ip

import (
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	schema "oduvanchiki/pkg/db"
	"os"
	"time"
)

type Storage struct {
	Db *gorm.DB
}

type ViewData struct {
	User      string
	Id        int
	Question  string
	Answer1   string
	Answer2   string
	Answer3   string
	Answer4   string
	TestStart int
	TestId    string
	Available bool
	Point     int
	Level     string
}

type FormData struct {
	Questionid string
	Answer     string
	TestStart  string
}

var errlog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var inflog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// Home Обработчик главной страницы.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера, запрашиваемая страница не найдена. %s", err)
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера. %s", err)
		http.Error(w, "Внутренняя ошибка сервера", 500)
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
		errlog.Printf("Внутренняя ошибка сервера, запрашиваемая страница не найдена. %s", err)
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера. %s", err)
		http.Error(w, "внутренняя ошибка сервера", 500)
	}
}

// NextTest Обработчик отображение страницы с формой
func (s *Storage) NextTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/next_test" {
		http.NotFound(w, r)
		return
	}

	user := schema.Clientusers{
		Name: r.FormValue("name"),
	}

	// Получить ip-адрес клиента
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Номер теста
	var numTest int

	if user.Name != "" {
		// Создать запись Clientusers
		recordName := &schema.Clientusers{Name: user.Name, Ip: ip}
		resultClient := s.Db.Create(&recordName)

		if resultClient.Error == nil {
			inflog.Printf("В Clientusers создано количество записей %v\n", resultClient.RowsAffected)
			inflog.Printf("Создана запись в Clientusers %v с его ip %v и получен id записи %v\n", recordName.Name, recordName.Ip, recordName.Id)

		} else {
			errlog.Printf("Не удалось создать запись имени %v\n", recordName.Name)
		}

		timeT := startTime()

		//Создать запись Quizes
		recordTest := schema.Quizes{Userid: recordName.Id, Started: timeT}
		resultQuiz := s.Db.Create(&recordTest)

		if resultQuiz.Error == nil {
			inflog.Printf("В Quizes создано количество записей %v\n", resultQuiz.RowsAffected)
			inflog.Printf("Создана запись в Quizes с временем %v и получен id %v\n", timeT, recordTest.Id)
			numTest = recordTest.Id
		} else {
			errlog.Printf("Не удалось создать запись теста %v\n", recordTest.Id)
		}

	} else {

		errlog.Print("Ошибка ввода имени")
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}

	data := ViewData{
		User:      user.Name,
		TestStart: numTest,
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/next_test.html")
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера, запрашиваемая страница не найдена. %v\n", err)
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, data)
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера. %s", err)
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

}

// FormTest Обработчик сохранения данных страницы с формой
func (s *Storage) FormTest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.NotFound(w, r)
		return
	}

	form := FormData{
		TestStart:  r.FormValue("id"),
		Answer:     r.FormValue("answer"),
		Questionid: r.FormValue("questionid"),
	}

	fmt.Println("/---------------------------------------------")

	// Извлечение объектов, где поле id равно form.TestStart
	var quizes schema.Quizes
	s.Db.Where("id = ?", form.TestStart).Find(&quizes)

	// Извлечение объектов, где поле answercorrect равно form.Answer
	var correct schema.Correctanswers
	s.Db.Where("answercorrect = ?", form.Answer).Find(&correct)
	inflog.Printf("Верный ответ = %v id = %v\n", correct.Correct, correct.Questionid)

	var result schema.Results
	// Правильный ответ
	if correct.Correct == true {
		result.Point = 1
	}

	if form.Questionid != "" {
		timeT := startTime()

		var inputQuestion schema.Quiestions
		// Извлечение объектов, где поле quiestionid равно form.Questionid
		s.Db.Where("id = ?", form.Questionid).Find(&inputQuestion)

		var inputAnswer schema.Answers
		// Извлечение объектов, где поле quiestionid равно form.Questionid
		s.Db.Where("quiestionid = ?", form.Questionid).Find(&inputAnswer)

		//Создать запись Results
		s.Db.Create(&schema.Results{Questionid: inputQuestion.Id, Answerid: inputAnswer.Id, Quizid: quizes.Id, Answered: timeT, Point: result.Point})
		inflog.Printf("Запись результата %v , %v , %v , %v , %v \n", inputQuestion.Id, inputAnswer.Id, quizes.Id, timeT, result.Point)
	}

	fmt.Println("/---------------------------------------------")

	var user schema.Clientusers
	var display ViewData

	var point []schema.Results
	s.Db.Where("quizid = ?", form.TestStart).Find(&point)

	if len(point) == 60 {
		display.Available = true

		point := testresult(point)
		result.Point = point
		inflog.Printf("Правильных ответов %v\n", result.Point)

		// Извлечение объектов, где поле id равно quizes.Userid
		s.Db.Where("id = ?", quizes.Userid).Find(&user)
		inflog.Printf("Имя : %v\n", user.Name)

		level := levelTest(result.Point)
		display.Level = level
		inflog.Printf("Знания равны : %v \n", display.Level)

	}

	var ress []schema.Results
	// Извлечение объектов, где поле quizid равно form.TestStart
	s.Db.Where("quizid = ?", form.TestStart).Find(&ress)

	var resFix schema.Results
	if form.Questionid != "" {
		// Извлечение объектов, где поле questionid равно form.Questionid
		s.Db.Where("questionid = ?", form.Questionid).Find(&resFix)
	}
	cheater := bagUpdateFix(ress, resFix.Questionid)
	inflog.Printf("Обновление страницы с вопросами, ЧИТ %v\n", cheater)

	if cheater == true {
		display.Available = true
	}
	inflog.Printf("Available %v\n", display.Available)

	fmt.Println("/---------------------------------------------")

	var question schema.Quiestions
	var answer schema.Answers

	// Извлечение всех объектов
	var allq []schema.Quiestions
	s.Db.Find(&allq)

	var resR []schema.Results
	// Извлечение объектов, где поле quizid равно form.TestStart
	s.Db.Where("quizid = ?", form.TestStart).Find(&resR)

	var strId int
	if len(point) <= 59 {
		var err error
		// Рандомно выбираем первичный ключ
		strId, err = randomId(allq, resR)
		if err != nil {
			panic(err)
		}
		inflog.Printf("Рандомно выбираем первичный ключ %v\n", strId)

		// Извлечение объекта с помощью первичного ключа
		s.Db.First(&question, strId)

		// Извлечение объектов, где поле quiestionid равно первичному ключу strId
		s.Db.Where("quiestionid = ?", strId).Find(&answer)
	}

	data := ViewData{
		Available: display.Available,
		User:      user.Name,
		Point:     result.Point,
		Level:     display.Level,
		TestId:    form.TestStart,
		Question:  question.Question,
		Id:        question.Id,
		Answer1:   answer.Answer1,
		Answer2:   answer.Answer2,
		Answer3:   answer.Answer3,
		Answer4:   answer.Answer4,
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	ts, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера, запрашиваемая страница не найдена. %v", err)
		http.Error(w, "Внутренняя ошибка сервера, запрашиваемая страница не найдена.", 500)
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, data)
	if err != nil {
		errlog.Printf("Внутренняя ошибка сервера. %v", err)
		http.Error(w, "Внутренняя ошибка сервера", 500)
	}

}

// Определяет есть такая запись или обновлена страница
func bagUpdateFix(ress []schema.Results, resFix int) bool {
	var fix bool

	sl := make([]int, 0, 60)
	for _, v := range ress {
		sl = append(sl, v.Questionid)
	}

	for i := 0; i < len(sl)-1; i++ {
		if sl[i] == resFix {
			fix = true
		}
	}

	return fix
}

// Подсчитывает уровень знаний по количеству ответов
func levelTest(point int) string {
	var ups string
	switch {
	case point <= 15:
		ups = "A1"
	case point <= 25:
		ups = "A2"
	case point <= 40:
		ups = "B1"
	case point <= 55:
		ups = "B2"
	case point > 55:
		ups = "C1"
	}
	return ups
}

// Считает количество правильных ответов
func testresult(point []schema.Results) int {
	var p int
	for _, v := range point {
		p += v.Point
	}

	return p
}

// Создает рандомно число
func randomId(allq []schema.Quiestions, resR []schema.Results) (int, error) {

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
	var shortest, longest *[]int
	// Меняем тип
	shortest = &slR
	longest = &slQ

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
	// Рандом Id
	rand.Shuffle(len(diff), func(i, j int) { diff[i], diff[j] = diff[j], diff[i] })

	return diff[0], nil
}

// Выводит время +Unix
func startTime() time.Time {

	tNow := time.Now()
	//Время для Unix Timestamp
	tUnix := tNow.Unix()
	//Временная метка Unix для time.Time
	time.Unix(tUnix, 0)

	return time.Now()
}
