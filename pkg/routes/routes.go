package routes

import (
	"example/api-call/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var todoList = []models.Todos{
// 	{ID: "1", Title: "clean room", Status: false},
// 	{ID: "2", Title: "make breakfast", Status: false},
// 	{ID: "3", Title: "exercise", Status: false},
// }

// func getTodos(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, todoList)
// }

func ping(context *gin.Context) {
	reqbody := "pong"
	context.IndentedJSON(http.StatusOK, reqbody)
}

// func postTodos(context *gin.Context) {
// 	var newTodo models.Todos
// 	err := context.BindJSON(&newTodo)
// 	if err != nil {
// 		return
// 	}
// 	todoList = append(todoList, newTodo)

// 	context.IndentedJSON(http.StatusOK, todoList)
// }
// func getTodosById(id string) (*models.Todos, error) {
// 	for i, t := range todoList {
// 		if t.ID == id {
// 			return &todoList[i], nil
// 		}
// 	}
// 	return nil, errors.New("not found by this id")
// }
// func getTodo(context *gin.Context) {
// 	id := context.Param("id")
// 	todo, err := getTodosById(id)
// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"messege": "to do not found"})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, todo)
// }

func register(context *gin.Context) {
	var newReq models.Todos
	err := context.BindJSON(&newReq)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	err = InsertTodo(newReq)
	if err != nil {
		log.Fatalf("error while inserting")
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successfully registered"})

}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// router.GET("/todos", getTodos)
	router.GET("/ping", ping)
	// router.POST("/todos", postTodos)
	// router.GET("/todos/:id", getTodo)
	router.POST("/add", register)
	log.Println("Running on port 8080")
	router.Run("localhost:8080")

}
