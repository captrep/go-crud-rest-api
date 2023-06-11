package app

import "github.com/captrep/go-crud-rest-api/controller"

func apiRoutes(userCtrl *controller.UserController) {
	r.Get("/ping", controller.PongHandler)
	r.Post("/api/v1/users", userCtrl.CreateUser)
	r.Get("/api/v1/users", userCtrl.FindAll)
	r.Get("/api/v1/users/{id}", userCtrl.FindById)
	r.Put("/api/v1/users/{id}", userCtrl.Update)
	r.Delete("/api/v1/users/{id}", userCtrl.Delete)
}
