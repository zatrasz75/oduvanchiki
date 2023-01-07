package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Quiestions struct {
	Id       int
	Question string
}

type Answer struct {
	Id          int
	Answer1     string
	Answer2     string
	Answer3     string
	Answer4     string
	QuiestionId string // Необходимо вставить Id из структуры Quiestions
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "rootroot"
	dbname   = "Dandelions" // Dandelions productdb
)

var (
	// Подключение к БД
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

// QuiestionDB Получение 1 записи из таблицы quiestions
func QuiestionOneDB(id int) Quiestions {
	// Открываем БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// Закрытие БД
	defer db.Close()
	// Контроль БД
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("select * from quiestions where id = $1", id)
	prod := Quiestions{}
	err = row.Scan(&prod.Id, &prod.Question)
	if err != nil {
		panic(err)
	}

	return prod
}

// SelectQuiestDB Получение всех записей из таблицы quiestions
func QuiestDB() []Quiestions {
	// Открываем БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// Закрытие БД
	defer db.Close()
	// Контроль БД
	err = db.Ping()

	rows, err := db.Query("select *from quiestions")
	if err != nil {
		panic(err)
	}

	data := []Quiestions{}

	for rows.Next() {
		p := Quiestions{}
		err := rows.Scan(&p.Id, &p.Question)
		if err != nil {
			panic(err)
		}
		data = append(data, p)
	}

	return data
}

// AnswersDB Получение всех записей из таблицы answer
func AnswersDB() []Answer {
	// Открываем БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// Закрытие БД
	defer db.Close()
	// Контроль БД
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select *from answer")
	if err != nil {
		panic(err)
	}

	data := []Answer{}

	for rows.Next() {
		p := Answer{}
		err := rows.Scan(&p.Id, &p.Answer1, &p.Answer2, &p.Answer3, &p.Answer4, &p.QuiestionId)
		if err != nil {
			panic(err)
		}
		data = append(data, p)
	}

	return data
}

// AnswerOneDB Получение 1 записи из таблицы answer
func AnswerOneDB(id int) Answer {
	// Открываем БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// Закрытие БД
	defer db.Close()
	// Контроль БД
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("select * from answer where id = $1", id)
	prod := Answer{}
	err = row.Scan(&prod.Id, &prod.Answer1, &prod.Answer2, &prod.Answer3, &prod.Answer4, &prod.QuiestionId)
	if err != nil {
		panic(err)
	}

	return prod
}
