package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"patricego/repositories/mysql"
	"patricego/services"
	"patricego/transport"
	"patricego/transport/endpoints"
	"patricego/usecases"
)
func main() {

	dsn := "root:Punsith12@tcp(127.0.0.1:3306)/taskDB?parseTime=true&loc=Asia%2FColombo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	fmt.Println(" Database connected successfully!")

	
	taskRepo := mysql.NewTaskRepository(db)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	taskService := services.NewTaskService(taskUsecase)
	taskHandler := endpoints.NewTaskHandler(taskService)

	router := transport.NewRouter(taskHandler)

	fmt.Println(" Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
