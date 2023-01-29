package db

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//	"log"
//)
//
//type client_user struct {
//	Name string
//}
//
//type Quiestions struct {
//	Id       int
//	Question string
//}
//
//type Answer struct {
//	Id          int
//	Answer1     string
//	Answer2     string
//	Answer3     string
//	Answer4     string
//	QuiestionId string // Необходимо вставить Id из структуры Quiestions
//}
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "rootroot"
//	dbname   = "Dandelions" // Dandelions productdb
//)
//
////======================================================================================================
//
////var dbase *gorm.DB
//
//var dsn = "host=localhost user=postgres password=rootroot dbname=Dandelions port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//
//func Init(n string) *gorm.DB {
//
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Нет подключения к БД", err)
//	}
//
//	err = db.AutoMigrate(&client_user{})
//	if err != nil {
//		return nil
//	}
//
//	db.Create(&client_user{Name: "получилось"})
//
//	return db
//}
//
////============================================================================================
//
//var (
//	// Подключение к БД
//	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//)
//
//// QuiestionOneDB QuiestionDB Получение 1 записи из таблицы quiestions
//func QuiestionOneDB(id int) Quiestions {
//	// Открываем БД
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	// Закрытие БД
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	// Контроль БД
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	row := db.QueryRow("select *from quiestions where id = $1", id)
//	prod := Quiestions{}
//	err = row.Scan(&prod.Id, &prod.Question)
//	if err != nil {
//		panic(err)
//	}
//
//	return prod
//}
//
//// QuiestDB Получение всех записей из таблицы quiestions
//func QuiestDB() []Quiestions {
//	// Открываем БД
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	// Закрытие БД
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	// Контроль БД
//	err = db.Ping()
//
//	rows, err := db.Query("select *from quiestions")
//	if err != nil {
//		panic(err)
//	}
//
//	var data []Quiestions
//
//	for rows.Next() {
//		p := Quiestions{}
//		err := rows.Scan(&p.Id, &p.Question)
//		if err != nil {
//			panic(err)
//		}
//		data = append(data, p)
//	}
//
//	return data
//}
//
//// AnswersDB Получение всех записей из таблицы answer
//func AnswersDB() []Answer {
//	// Открываем БД
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	// Закрытие БД
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	// Контроль БД
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	rows, err := db.Query("select *from answer")
//	if err != nil {
//		panic(err)
//	}
//
//	var data []Answer
//
//	for rows.Next() {
//		p := Answer{}
//		err := rows.Scan(&p.Id, &p.Answer1, &p.Answer2, &p.Answer3, &p.Answer4, &p.QuiestionId)
//		if err != nil {
//			panic(err)
//		}
//		data = append(data, p)
//	}
//
//	return data
//}
//
//// AnswerOneDB Получение 1 записи из таблицы answer
//func AnswerOneDB(id int) Answer {
//	// Открываем БД
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	// Закрытие БД
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	// Контроль БД
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	row := db.QueryRow("select * from answer where id = $1", id)
//	prod := Answer{}
//	err = row.Scan(&prod.Id, &prod.Answer1, &prod.Answer2, &prod.Answer3, &prod.Answer4, &prod.QuiestionId)
//	if err != nil {
//		panic(err)
//	}
//
//	return prod
//}
//
//func NameAnswerDB(name string) int {
//	// Открываем БД
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	// Закрытие БД
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	// Контроль БД
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	var id int
//	fmt.Println(name)
//	db.QueryRow("insert into 'client_user' ('user_name')values ('go-20')returning id").Scan(id)
//
//	return id
//}
