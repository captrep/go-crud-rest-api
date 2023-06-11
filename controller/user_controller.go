package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/captrep/go-crud-rest-api/helper"
	"github.com/captrep/go-crud-rest-api/model/web"
	"github.com/captrep/go-crud-rest-api/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUserRequest := web.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		panic(err)
	}
	userResponse, err := ctrl.svc.CreateUser(&createUserRequest)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			helper.ValidationError(w, err.(validator.ValidationErrors))
		} else {
			helper.InternalServerError(w, err.Error())
		}

		panic(err)
	}
	webResponse := &web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   userResponse,
	}
	helper.WriteJSON(w, http.StatusCreated, webResponse)
}

func (ctrl *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	userResponses, err := ctrl.svc.FindAll()
	if err != nil {
		helper.InternalServerError(w, err.Error())
		log.Println(err.Error())
		return
	}
	resp := &web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponses,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (ctrl *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	userResponse, err := ctrl.svc.FindById(userId)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			helper.NotFoundError(w, err.Error())
		} else {
			helper.InternalServerError(w, err.Error())
		}
		return
	}
	resp := &web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (ctrl *UserController) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	updateUserRequest := web.UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&updateUserRequest)
	if err != nil {
		panic(err)
	}
	updateUserRequest.Id = userId
	res, err := ctrl.svc.Update(&updateUserRequest)
	if err != nil {
		log.Println(err.Error())
		log.Println(err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			helper.NotFoundError(w, err.Error())
		} else {
			helper.InternalServerError(w, err.Error())
		}
		return
	}
	webResponse := &web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   res,
	}

	helper.WriteJSON(w, http.StatusOK, webResponse)
}

func (ctrl *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	res, err := ctrl.svc.Delete(userId)
	if err != nil {
		log.Println(err.Error())
		log.Println(err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			helper.NotFoundError(w, err.Error())
		} else {
			helper.InternalServerError(w, err.Error())
		}
		return
	}
	helper.WriteJSON(w, http.StatusOK, res)
}
