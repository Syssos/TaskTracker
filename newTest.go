package main

import ()

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

type DatabaseCoupler struct {
	Projects []project
}

var AppProjectsInMem DatabaseCoupler{}

func (db *DatabaseCoupler) Load() {

}

func (db *DatabaseCoupler) Commit() {

}

func (db *DatabaseCoupler) Renew() {
	
}