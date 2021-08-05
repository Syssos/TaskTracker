package project

import (
	"fmt"
	"time"
	"net/http"
	"math/rand"
	"encoding/json"
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

func NewProject(r *http.Request) *Project {
	var pro Project
	pro.ID = String(6)
	err := json.NewDecoder(r.Body).Decode(&pro)
	if err != nil {
		fmt.Println(err)
		return &Project{}
	}

	return &pro
}

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
