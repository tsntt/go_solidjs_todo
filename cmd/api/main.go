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

	taskHandlers.Create("take a shower", "wash face", "1695949414000")
	taskHandlers.Create("take another shower", "wash feet", "1698541414000")

	server := api.NewApiServer(taskHandlers)

	server.Run()

}
