package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	data "tsn/todo/data/postgres"
	"tsn/todo/src/entities"
	"tsn/todo/src/usecases"
	"tsn/todo/src/util"
)

func main() {
	whichDB := flag.String("db", "postgres", "select database")
	flag.Parse()

	fmt.Printf("%s\n", *whichDB)

	var envVars map[string]string
	var store entities.TaskRepository

	err := util.SetEnvs("env.json", &envVars)
	if err != nil {
		log.Fatal(err)
	}

	if *whichDB == "postgres" {
		store, err = data.NewPostgresStore(envVars["DBCONN"])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("you shoud chose 'postgres' or implement a new db")
		return
	}

	jsonFile, err := os.Open("./cmd/seed/seeds.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}

	var tasks []map[string]interface{}

	if err = json.Unmarshal(byteValue, &tasks); err != nil {
		log.Fatal(err)
	}

	taskHandlers := usecases.NewTaskInteractor(store)

	for _, task := range tasks {
		_, err = taskHandlers.Create(fmt.Sprintf("%v", task["content"]), fmt.Sprintf("%v", task["description"]), fmt.Sprintf("%v", task["due"]))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("DB seed with %d entries\n", len(tasks))
}
