package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Task1", "Task2"}

func main() {
	fmt.Println("#### Welcome to our todo list! ####")

	http.HandleFunc("/", helloUser) // this is a handler, first param is a url

	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe("localhost:8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello World"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}
