package routes

import (
	"fmt"
  "time"
  "path"
	"net/http"
	"math/rand"
	"encoding/json"

	"github.com/gorilla/mux"
)

// ****************************************************************************************************************************************
// Project Structures Used for tracking notes.
// This project will deal with reading and writing tasks to a file for long term storage

// AS OF RIGHT NOW THE PROJECT STORES EVERYTHING IN MEMORY TO AVOID KEEPING TRACK OF THE DATA FILE OR DATA STORED IN A FILE


type Task struct {
	ID        string `json: "id"`
	TaskTitle string `json: "tasktitle"`
	TaskData  string `json: "taskdata"`
}

type Project struct {
	ID            string `json: "id"`
	ProjectTitle  string `json: "projecttitle"`
	Tasks         []Task `json: "Tasks"`
}

// ****************************************************************************************************************************************
// Createing Temp in memory storage for project to use while testing

var App_Instance_Projects = []Project{
	{ID: "L62L0p", ProjectTitle: "Perminant Task", Tasks: []Task{Task{ID: "CJ1a32", TaskTitle: "Task 1", TaskData: "This is a test task"},Task{ID: "xs4oih", TaskTitle: "Task 2", TaskData: "This is a test task"}}},
}

func projectFromAppList(id string) *Project {
	for _, item := range App_Instance_Projects {
		if item.ID == id {
			return &item
		}
	}
	return &Project{}
}

func updateProject(pjk *Project) {
	for x, item := range App_Instance_Projects {
		if item.ID == pjk.ID {
			App_Instance_Projects[x] = *pjk
		}
	}
}

// ****************************************************************************************************************************************
// Routes used by the backend to send and receive data

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%v", App_Instance_Projects)
}

func NewProject(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/new/project" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var pro Project
	pro.ID = String(6)
	err := json.NewDecoder(r.Body).Decode(&pro)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	App_Instance_Projects = append(App_Instance_Projects, pro)

	payload, _ := json.Marshal(App_Instance_Projects)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for x, item := range App_Instance_Projects {
		if string(item.ID) == vars["proId"] {
			var pro Project
			err := json.NewDecoder(r.Body).Decode(&pro)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			prj := projectFromAppList(item.ID)
			prj.ProjectTitle = pro.ProjectTitle
			fmt.Println(prj)
			App_Instance_Projects[x] = *prj
			// prj.TaskData = tsk.TaskData
		}
	}
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, item := range App_Instance_Projects {
		if string(item.ID) == vars["proId"] {
			var tsk Task
			tsk.ID = String(6)
			err := json.NewDecoder(r.Body).Decode(&tsk)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			pjk := projectFromAppList(item.ID)
			pjk.Tasks = append(pjk.Tasks, tsk)
			updateProject(pjk)

			payload, _ := json.Marshal(pjk.Tasks)

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(payload))
		}
	}
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	payload, _ := json.Marshal(App_Instance_Projects)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	pro := path.Base(r.URL.Path)
	url := r.URL.Path[:len(r.URL.Path)-len(pro)-1]
	if url != "/project" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	prj := projectFromAppList(pro)

	payload, _ := json.Marshal(prj)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

// *****************************************************************************************************************

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}

