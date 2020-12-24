package main

import (
	"encoding/json"
	"errors"
	"example/api"
	"example/dao"
	"example/status"
	"example/task"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Response is JSON response object
type Response struct {
	ID    dao.ID            `json:"id,string"`
	Task  task.Task         `json:"task,"`
	Error api.ResponseError `json:"error,"`
}

const (
	pathPrefix = "/api/v1/task/"
	htmlPrefix = "/task/"
)

// NOT thread-safe
var m = dao.NewMemoryDAO()

var tmpl = template.Must(template.ParseGlob("html/*.html"))

func apiHandler(w http.ResponseWriter, req *http.Request) {
	getID := func() (dao.ID, error) {
		id := dao.ID(req.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, fmt.Errorf("main.apiHandler: id is empty; %s", req.URL.Path)
		}
		return id, nil
	}
	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		if err := req.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := req.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t task.Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}
	switch req.Method {
	case http.MethodGet:
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: api.ResponseError{Err: err},
		})
		if err != nil {
			log.Println(err)
		}
	case http.MethodPut:
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: api.ResponseError{Err: err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case http.MethodPost:
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: api.ResponseError{Err: err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case http.MethodDelete:
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Error: api.ResponseError{Err: err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func htmlHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Println("unsupported method", req.Method)
		return
	}
	getID := func() (dao.ID, error) {
		id := dao.ID(req.URL.Path[len(htmlPrefix):])
		if id == "" {
			return id, fmt.Errorf("htmlHandler: id is empty; %s", req.URL.Path)
		}
		return id, nil
	}
	id, err := getID()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := m.Get(id)
	err = tmpl.ExecuteTemplate(w, "task.html", &Response{
		ID:    id,
		Task:  t,
		Error: api.ResponseError{Err: err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func init() {
	// Database fixture
	id, _ := m.Post(task.Task{Title: "Laundry", Status: status.WIP})
	fmt.Println(id)
	id, _ = m.Post(task.Task{Title: "Greeting a friend", Status: status.WIP})
	fmt.Println(id)
	id, _ = m.Post(task.Task{Title: "Study at library", Status: status.Done})
	fmt.Println(id)
}

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	http.HandleFunc(htmlPrefix, htmlHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
