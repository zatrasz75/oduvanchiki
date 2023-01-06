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

// RowDB Получение данных из БД по 1 записи
func RowDB(id int) Quiestions {
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

// SelectDB Получение данных из БД
func SelectDB() []Quiestions {
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
