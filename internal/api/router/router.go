package router

import (
	"conn/internal/api/handler"
	"conn/internal/database"
	"conn/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	r := http.NewServeMux()

	psql, err := connection()
	if err != nil {
		log.Fatal(err)
	}
	postgres := database.NewPostgres(psql)

	service := service.NewService(postgres)

	handler := handler.NewHandler(service)

	r.HandleFunc("GET /books/{id}", handler.GetBook)
	r.HandleFunc("GET /books", handler.GetBooks)
	r.HandleFunc("POST /books", handler.BookCreate)
	r.HandleFunc("PUT /books/{id}", handler.BookUpdate)
	r.HandleFunc("DELETE /books/{id}", handler.DeleteBook)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:2005@mypost2:5432/docker?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
