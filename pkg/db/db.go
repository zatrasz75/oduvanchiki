package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "759704"
	dbname   = "productdb"
)

var (
	// Подключение к БД
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func СonnectDB() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

// RowDB Получение данных из БД по 1 записи
func RowDB(id int) Product {
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

	row := db.QueryRow("select * from Products where id = $1", id)
	prod := Product{}
	err = row.Scan(&prod.Id, &prod.Model, &prod.Company, &prod.Price)
	if err != nil {
		panic(err)
	}

	return prod
}
