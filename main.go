package main

import (
	"log"
	"net/http"
	"todo-api/controllers"
	_ "todo-api/docs"
	"todo-api/models"
	"todo-api/repositories"
	"todo-api/services"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title To-Do API
// @version 1.0
// @description This is a simple MVC Api made by Jonathan Malagueta
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	dsn := "host=localhost user=example_user password=example_password dbname=go_basic port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Dropar a tabela existente (apenas para desenvolvimento)
	db.Migrator().DropTable(&models.Task{})

	// Migrar o banco de dados para criar a tabela com as configurações corretas
	db.AutoMigrate(&models.Task{})

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskController.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskController.DeleteTask).Methods("DELETE")

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
