package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"tsn/todo/src/usecases"
	"tsn/todo/src/util"
)

type ApiServer struct {
	services *usecases.TaskInteractor
}

func NewApiServer(services *usecases.TaskInteractor) *ApiServer {
	return &ApiServer{
		services: services,
	}
}

func (s *ApiServer) Run() {
	routes := http.NewServeMux()

	solidjs := http.FileServer(http.Dir("./view/solidjs"))

	routes.Handle("/", solidjs)
	routes.HandleFunc("/getall", s.HandlerIndex)
	routes.HandleFunc("/create", s.HandlerCreate)
	routes.HandleFunc("/update/", s.HandlerUpdate)
	routes.HandleFunc("/changestatus/", s.HandlerChangeStatus)
	routes.HandleFunc("/delete/", s.HandlerDelete)

	log.Println("Api listening at port: 4000")
	log.Fatal(http.ListenAndServe(":4000", routes))
}

func (s *ApiServer) HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteJson(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}
	tasks, err := s.services.GetAll()
	if err != nil {
		util.WriteJson(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	util.WriteJson(w, http.StatusOK, tasks)
}

func (s *ApiServer) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		util.WriteJson(w, http.StatusOK, "")
		return
	}
	if r.Method != "POST" {
		util.WriteJson(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	vars, err := util.ReadJson(w, r)
	if err != nil {
		util.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	newTask, err := s.services.Create(vars["content"], vars["due"])
	if err != nil {
		util.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.WriteJson(w, http.StatusOK, newTask)
}

func (s *ApiServer) HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		util.WriteJson(w, http.StatusOK, "")
		return
	}
	if r.Method != "PUT" {
		util.WriteJson(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/update/"))
	if err != nil {
		util.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	vars, err := util.ReadJson(w, r)
	if err != nil {
		util.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	task, err := s.services.Update(id, vars["content"], vars["due"])
	if err != nil {
		util.WriteJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	util.WriteJson(w, http.StatusOK, task)
}

func (s *ApiServer) HandlerChangeStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		util.WriteJson(w, http.StatusOK, "")
		return
	}
	if r.Method != "PATCH" {
		util.WriteJson(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/changestatus/"))
	if err != nil {
		util.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	task, err := s.services.ChangeStatus(id)
	if err != nil {
		util.WriteJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	util.WriteJson(w, http.StatusOK, task)
}

func (s *ApiServer) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		util.WriteJson(w, http.StatusOK, "")
		return
	}
	if r.Method != "DELETE" {
		util.WriteJson(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/delete/"))
	if err != nil {
		util.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = s.services.Delete(id)
	if err != nil {
		util.WriteJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	util.WriteJson(w, http.StatusOK, map[string]int{"deleted": id})
}
