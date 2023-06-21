package helper

import (
	"net/http"

	"github.com/captrep/go-crud-rest-api/model/web"
)

func ValidationError(w http.ResponseWriter, data interface{}) {
	WriteJSON(w, http.StatusBadRequest, &web.ErrorResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
		Error:  data,
	})
}

func NotFoundError(w http.ResponseWriter, err string) {
	WriteJSON(w, http.StatusNotFound, &web.ErrorResponse{
		Code:   http.StatusNotFound,
		Status: "NOT FOUND",
		Error:  err,
	})
}

func InternalServerError(w http.ResponseWriter, err string) {
	WriteJSON(w, http.StatusInternalServerError, &web.ErrorResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err,
	})
}
