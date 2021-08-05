package routes

import (
	"fmt"
  "time"
  "path"
	"net/http"
	"math/rand"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.comSyssos/ProjectTaskTracker/pkg/project"
)

// ****************************************************************************************************************************************
// Createing Temp in memory storage for project to use while testing

var App_Instance_Projects = []project.Project{
	{ID: "L62L0p", ProjectTitle: "Perminant Task", Tasks: []project.Task{{ID: "CJ1a32", TaskTitle: "Task 1", TaskData: "This is a test task"},{ID: "xs4oih", TaskTitle: "Task 2", TaskData: "This is a test task"}}},
}

func projectFromAppList(id string) *project.Project {
	for _, item := range App_Instance_Projects {
		if item.ID == id {
			return &item
		}
	}
	return &project.Project{}
}

func updateProject(pjk *project.Project) {
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

	pro := project.NewProject(r)
	App_Instance_Projects = append(App_Instance_Projects, *pro)
	payload, _ := json.Marshal(App_Instance_Projects)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for x, item := range App_Instance_Projects {
		if string(item.ID) == vars["proId"] {
			var pro project.Project
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
			var tsk project.Task
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

