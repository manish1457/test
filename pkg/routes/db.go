package routes

import (
	"database/sql"
	"example/api-call/pkg/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	dbname   = "testdb"
	user     = "usermanish"
	password = "mypass"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("from here")
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("i am here")
		panic(err)
	}
	log.Println("db connected")
	return db
}

func InsertTodo(todo models.Todos) error {
	db := InitDB()
	sqlStatement := `INSERT INTO todo_list_table (id, title, status) VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, todo.ID, todo.Title, todo.Status)
	if err != nil {
		log.Printf("hello error:%v", err)
		return err
	}
	return nil
}
