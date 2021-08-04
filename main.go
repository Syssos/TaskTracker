package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.comSyssos/ProjectTaskTracker/pkg/routes"
)

func setRoutes(){
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home)
	r.HandleFunc("/projects", routes.GetProjects)
	r.HandleFunc("/project/{pro}", routes.GetProject)
	r.HandleFunc("/new/project", routes.NewProject)
	r.HandleFunc("/update/project/{proId}", routes.UpdateProject)
	r.HandleFunc("/new/task/{proId}", routes.NewTask)

	r.NotFoundHandler = http.HandlerFunc(notFound)

	fmt.Println("Starting Server on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	return
}

func main() {
	setRoutes()
}