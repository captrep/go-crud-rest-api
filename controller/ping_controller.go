package controller

import (
	"encoding/json"
	"net/http"

	"github.com/captrep/go-crud-rest-api/model/web"
)

func PongHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp := &web.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   "pong",
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
