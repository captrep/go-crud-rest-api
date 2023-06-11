package app

import (
	"log"
	"net/http"

	"github.com/captrep/go-crud-rest-api/controller"
	"github.com/captrep/go-crud-rest-api/datasource/mysql"
	"github.com/captrep/go-crud-rest-api/repository"
	"github.com/captrep/go-crud-rest-api/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

var r = chi.NewRouter()

func StartApplication() {
	db := mysql.NewDB()
	validate := validator.New()
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, validate)
	userCtrl := controller.NewUserController(userSvc)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	log.Println("Server running on port 8k")
	apiRoutes(userCtrl)
	http.ListenAndServe("127.0.0.1:8000", r)
}
