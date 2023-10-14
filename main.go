package main

import (
	"log"
	"os"
	"strings"
)

type Todo struct {
	title string
	desc  string
}

func main() {
	if len(os.Args) <= 1 {
		printUsage()
		return
	}

	fileName := "todos.txt"
	todos := getTodos(fileName)

	switch os.Args[1] {
	case "create":
		if len(os.Args) < 4 {
			printUsage()
			return
		}
		todos = append(todos, Todo{os.Args[2], os.Args[3]})
		println("Created:" + os.Args[2])
		writeTodosToFile(todos, fileName)
		return
	case "list":
		printTodo(todos)
		return
	case "delete":
		if len(os.Args) < 3 {
			printUsage()
			return
		}
		todos = removeTodo(todos, os.Args[2])
		writeTodosToFile(todos, fileName)
	default:
		printUsage()
	}
}

func writeTodosToFile(todos []Todo, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for t := range todos {
		_, err := f.WriteString(todos[t].title + "\n" + todos[t].desc + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func printTodo(todos []Todo) {
	for t := range todos {
		println(":" + todos[t].title + ":")
		if len(todos) <= 1 {
			println(todos[t].desc)
			continue
		}
		println(todos[t].desc, "\n")
	}
}

func getTodos(filename string) []Todo {
	t := make([]Todo, 0)
	ts, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(ts), "\n")

	for i := 0; i < len(lines)-1; i += 2 {
		t = append(t, Todo{lines[i], lines[i+1]})
	}
	return t
}

func removeTodo(todos []Todo, title string) []Todo {
	t := make([]Todo, 0)
	for i := 0; i < len(todos); i++ {
		if todos[i].title != title {
			t = append(t, todos[i])
		} else {
			println("Deleted:" + todos[i].title)
		}
	}
	return t
}

func printUsage() {
	println("Usage: create <title> <description>")
	println("Usage: list")
	println("Usage: delete <title>")
}
