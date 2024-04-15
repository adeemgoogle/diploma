package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func InitDB() {
	var err error
	// Открываем соединение с базой данных PostgreSQL
	db, err = sqlx.Open("postgres", "user=postgres password=123 dbname=Diploma sslmode=disable")
	if err != nil {
		log.Fatalln("Ошибка при подключении к базе данных:", err)
	}
	// Проверяем соединение с базой данных
	if err = db.Ping(); err != nil {
		log.Fatalln("Ошибка при проверке соединения с базой данных:", err)
	}
	log.Println("Подключение к базе данных успешно установлено")
}
