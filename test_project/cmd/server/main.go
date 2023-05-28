package main

import (
	"flag"
	"fmt"
	"github.com/login/test_project/internal/use_case/repository/postgres"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/login/test_project/internal/service/friend_service"
	"github.com/login/test_project/internal/use_case/friend"
	"github.com/login/test_project/pkg/db"
)

func main() {
	port := flag.String("port", ":3000", "is port valuese")
	flag.Parse()

	// создаем экземплям бд
	postgresDB, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	// создаем репозиторий для хранения данных (map)
	repository := postgres.NewPostgresRepository(postgresDB)

	// создаем use_case
	friendUseCase := friend.New(repository)

	// создаем сервис для обработки запросов клиента
	friendService := friend_service.New(friendUseCase)

	// создаем новый роутер
	r := chi.NewRouter()

	// показывает информацию по запросам
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, i'm a golang microservice")
		w.WriteHeader(http.StatusAccepted)
	})

	// создать друга
	r.Post("/create", friendService.CreateFriend)
	// подружить двух друзей
	r.Post("/make_friends", friendService.MakeFriend)
	// показать всех друзей пользователя
	r.Get("/friends/{user_id}", friendService.GetFriends)
	// удалить пользователя
	r.Delete("/{user}", friendService.DeleteFriend)
	// Обновить возраст пользователя
	r.Put("/{user_id}", friendService.UpdateAge)

	// запускаем работу сервиса на 8080 порту
	log.Printf("service start %s port", *port)
	if err := http.ListenAndServe(*port, r); err != nil {
		log.Fatal(err)
	}
}
