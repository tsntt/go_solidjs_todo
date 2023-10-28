package main

import (
	"tsn/todo/api"
	data "tsn/todo/data/memory"
	"tsn/todo/src/entities"
	"tsn/todo/src/usecases"
)

func main() {

	var tasklist []entities.Task

	db := data.NewMemoryStorage(tasklist)

	taskHandlers := usecases.NewTaskInteractor(db)

	taskHandlers.Create("take a shower", "1695949414")
	taskHandlers.Create("take another shower", "1698541414")

	server := api.NewApiServer(taskHandlers)

	server.Run()

}
