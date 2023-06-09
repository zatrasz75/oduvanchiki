package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	schema "oduvanchiki/pkg/db"
	"oduvanchiki/pkg/ip"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

const (
	Host     = "localhost"
	Port     = 5432
	Users    = "postgres"
	Password = "root"
	Dbname   = "Dandelions"
)

var errlog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var inflog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

var (
	// Подключение к БД
	conStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, Users, Password, Dbname)
)

func main() {
	// Соединение с БД.
	s, err := ip.New(conStr)
	if err != nil {
		panic("не удалось подключить базу данных")
	}
	if err != nil {
		errlog.Fatal("Нет подключения к БД \n", err.Error())
	}

	// Удалить таблицы, если они существуют
	err = s.Db.Migrator().DropTable(&schema.Results{}, &schema.Quizes{}, &schema.Clientusers{})
	if err != nil {
		errlog.Printf("Не удалось удалить таблицы", err)
		return
	}

	// Удалить таблицы, если они существуют
	err = s.Db.Migrator().DropTable(&schema.Quiestions{}, &schema.Answers{}, &schema.Correctanswers{})
	if err != nil {
		errlog.Printf("Не удалось удалить таблицы %v", err)
	}

	// err = s.Db.Migrator().DropTable(&schema.AccountMail{}) //==========================================
	// if err != nil {
	// 	errlog.Printf("Не удалось удалить таблицы", err)
	// 	return
	// }

	// Перенос схемы в таблицу
	err = s.Db.AutoMigrate(&schema.Quiestions{}, &schema.Correctanswers{}, &schema.Answers{})
	if err != nil {
		errlog.Printf("Не удалось перенести схему %v", err)
	}

	// Добавление записи Quiestions
	pasteQuiestions, err := questionPaste()
	if err != nil {
		return
	}
	var question schema.Quiestions
	for i, v := range pasteQuiestions {
		if i == 0 {
			continue
		}
		question.Id = v.Id
		question.Question = v.Question
		result := s.Db.Create(&question)
		inflog.Printf("Создана %v запись Quiestions :\n %v\n", result.RowsAffected, v.Question)
	}

	// Добавление записи Correctanswers
	pasteCorrect, err := correctanswersPaste()
	if err != nil {
		return
	}
	var correct schema.Correctanswers
	for i, v := range pasteCorrect {
		if i == 0 {
			continue
		}
		correct.Id = v.Id
		correct.Questionid = v.Questionid
		correct.Answercorrect = v.Answercorrect
		result := s.Db.Create(&correct)
		inflog.Printf("Создана %v запись Correctanswers :\n %v\n", result.RowsAffected, v.Answercorrect)
	}

	// Добавление записи Answers
	pasteAnswer, err := answerPaste()
	if err != nil {
		return
	}
	var answer schema.Answers
	for i, v := range pasteAnswer {
		if i == 0 {
			continue
		}
		answer.Id = v.Id
		answer.Answer1 = v.Answer1
		answer.Answer2 = v.Answer2
		answer.Answer3 = v.Answer3
		answer.Answer4 = v.Answer4
		answer.Quiestionid = v.Quiestionid
		result := s.Db.Create(&answer)
		inflog.Printf("Создана %v запись Answers :\n %v, %v, %v, %v\n", result.RowsAffected, v.Answer1, v.Answer2, v.Answer3, v.Answer4)
	}

	// Перенос схемы в таблицу
	err = s.Db.AutoMigrate(&schema.Clientusers{}, &schema.Quizes{}, &schema.Results{})
	if err != nil {
		errlog.Printf("Не удалось перенести схему %v", err)
	}

	// Перенос схемы в таблицу
	err = s.Db.AutoMigrate(&schema.AccountMail{})

	//---------------------------------------------------------

	PORT := ":4000"

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./static".
	var fs = http.FileServer(http.Dir("./static/"))

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	srv := &http.Server{
		Handler:      router,
		Addr:         PORT,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	router.HandleFunc("/", ip.Home).Methods("GET")
	router.HandleFunc("/next_test", s.NextTest).Methods("POST")
	router.HandleFunc("/test", s.FormTest).Methods("POST")
	router.HandleFunc("/info-customer", ip.Customer).Methods("GET")
	router.HandleFunc("/connection", s.Connect).Methods("POST")

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

// answerPaste выводит данные из файла в массив.
func answerPaste() ([]schema.Answers, error) {
	f, err := os.Open("./pkg/db/answers_202302122152.csv")
	if err != nil {
		return nil, err
	}
	var slStr []schema.Answers
	reader := csv.NewReader(f)
	for {
		var str schema.Answers
		records, err := reader.Read()
		if err != nil {
			break
		}
		pId, _ := strconv.Atoi(records[0])
		str.Id = pId
		str.Answer1 = records[1]
		str.Answer2 = records[2]
		str.Answer3 = records[3]
		str.Answer4 = records[4]
		pQuiestionid, _ := strconv.Atoi(records[5])
		str.Quiestionid = pQuiestionid
		slStr = append(slStr, str)
	}

	return slStr, nil
}

// questionPaste выводит данные из файла в массив.
func questionPaste() ([]schema.Quiestions, error) {
	f, err := os.Open("./pkg/db/quiestions_202302122153.csv")
	if err != nil {
		return nil, err
	}
	var slStr []schema.Quiestions
	reader := csv.NewReader(f)
	for {
		var str schema.Quiestions
		records, err := reader.Read()
		if err != nil {
			break
		}
		pId, _ := strconv.Atoi(records[0])
		str.Id = pId
		str.Question = records[1]
		slStr = append(slStr, str)

	}

	return slStr, nil
}

// correctanswersPaste выводит данные из файла в массив.
func correctanswersPaste() ([]schema.Correctanswers, error) {
	f, err := os.Open("./pkg/db/correctanswers_202302142246.csv")
	if err != nil {
		return nil, err
	}
	var slStr []schema.Correctanswers
	reader := csv.NewReader(f)
	for {
		var str schema.Correctanswers
		records, err := reader.Read()
		if err != nil {
			break
		}
		pId, _ := strconv.Atoi(records[0])
		str.Id = pId
		pQuestionid, _ := strconv.Atoi(records[1])
		str.Questionid = pQuestionid
		//fmt.Println(records[0], records[1], records[2])
		str.Answercorrect = records[2]
		slStr = append(slStr, str)

	}

	return slStr, nil
}
