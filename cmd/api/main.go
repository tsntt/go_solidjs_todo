package main

import (
	"log"
	"tsn/todo/api"
	mem "tsn/todo/data/memory"
	data "tsn/todo/data/postgres"
	"tsn/todo/src/entities"
	"tsn/todo/src/usecases"
	"tsn/todo/src/util"
)

func main() {
	var db entities.TaskRepository
	var envVars map[string]string

	err := util.SetEnvs("env.json", &envVars)
	if err != nil {
		log.Fatal(err)
	}

	if envVars["USEDB"] == "memory" {
		var tasklist []entities.Task
		db = mem.NewMemoryStorage(tasklist)
	}

	if envVars["USEDB"] == "postgres" {
		db, err = data.NewPostgresStore(envVars["DBCONN"])
		if err != nil {
			log.Fatal(err)
		}
	}

	taskHandlers := usecases.NewTaskInteractor(db)
	server := api.NewApiServer(taskHandlers)

	server.Run()
}
