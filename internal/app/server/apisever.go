package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/albakov/go-todo/internal/app/model"
	"github.com/albakov/go-todo/internal/app/store"
	"github.com/gorilla/mux"
)

type apiserver struct {
	router *mux.Router
	store  store.Store
}

func apiServer() *apiserver {
	s := &apiserver{
		router: mux.NewRouter(),
		store:  store.Store{},
	}
	s.setRoutes()

	return s
}

func (s *apiserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *apiserver) respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *apiserver) setRoutes() {
	sub := s.router.PathPrefix("/api").Subrouter()
	sub.HandleFunc("/items", s.getItems).Methods("GET")
	sub.HandleFunc("/items", s.createItem).Methods("POST")
	sub.HandleFunc("/items/{id:[0-9]+}", s.getItem).Methods("GET")
	sub.HandleFunc("/items/{id:[0-9]+}", s.updateItem).Methods("POST")
	sub.HandleFunc("/items/{id:[0-9]+}", s.deleteItem).Methods("DELETE")
}

func (s *apiserver) getItems(w http.ResponseWriter, r *http.Request) {
	items := s.store.Item().Get()
	s.respond(w, http.StatusOK, items)
}

func (s *apiserver) getItem(w http.ResponseWriter, r *http.Request) {
	param, found := mux.Vars(r)["id"]
	if !found {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	item, err := s.store.Item().GetById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.respond(w, http.StatusOK, item)
}

func (s *apiserver) createItem(w http.ResponseWriter, r *http.Request) {
	data := &model.Item{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	item, err := s.store.Item().Create(*data)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.respond(w, http.StatusAccepted, item)
}

func (s *apiserver) updateItem(w http.ResponseWriter, r *http.Request) {
	param, found := mux.Vars(r)["id"]
	if !found {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := &model.Item{}

	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = s.store.Item().Update(id, *data)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.respond(w, http.StatusAccepted, nil)
}

func (s *apiserver) deleteItem(w http.ResponseWriter, r *http.Request) {
	param, found := mux.Vars(r)["id"]
	if !found {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = s.store.Item().Delete(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.respond(w, http.StatusAccepted, nil)
}
