package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"udon-web/models"
	"udon-web/repos"

	"github.com/gorilla/mux"
)

var udonRepo = repos.NewUdonRepo()

func CreateUdon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var udon models.Udon
	json.NewDecoder(r.Body).Decode(&udon)
	udon = udonRepo.Create(udon)
	json.NewEncoder(w).Encode(udon)
	w.WriteHeader(http.StatusCreated)
}

func GetUdons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(udonRepo.GetList())
	w.WriteHeader(http.StatusOK)
}

func GetUdonById(w http.ResponseWriter, r *http.Request) {
	udonIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	udon, err := udonRepo.GetOne(uint(udonIdLong))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Udon not found!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(udon)
	w.WriteHeader(http.StatusOK)
}

func UpdateUdon(w http.ResponseWriter, r *http.Request) {
	udonIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	var udon models.Udon
	json.NewDecoder(r.Body).Decode(&udon)
	_, err = udonRepo.Update(uint(udonIdLong), udon)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Udon not found!")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUdon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	udonIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	_, err = udonRepo.DeleteOne(uint(udonIdLong))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Udon not found!")
		return
	}

	json.NewEncoder(w).Encode("Udon Deleted Successfully!")
	w.WriteHeader(http.StatusNoContent)
}