package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Json200Response(w http.ResponseWriter, x interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(x)
}

func Json201Response(w http.ResponseWriter, x interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(x)
}

func Json400Response(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func Json404Response(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func GetIdParam(w http.ResponseWriter, r *http.Request) int64 {
	var id int64
	if intId, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0); err != nil {
		Json400Response(w, err)
		return 0
	} else {
		id = int64(intId)
	}
	return id
}

func ParseBody(w http.ResponseWriter, r *http.Request, x interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		Json400Response(w, err)
		return false
	}
	return true
}
