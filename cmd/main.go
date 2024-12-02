package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"medods/internal/handlers"
	"medods/internal/repository"
)

func main() {
	var port = "8080"
	// метод создаёт коннект с базой данных по паттерну репозиторий
	rep := repository.CreateRepository()
	// функция для подготовки таблицы, хранящей bcrypt-хэши
	err := rep.PrepareTable()
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(rep)

	router := chi.NewRouter()
	// хэндлер для создании пары RT-AT
	router.Get("/sign", handler.Sign)
	// хэндлер для обновления AT
	router.Get("/refresh", handler.Refresh)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
