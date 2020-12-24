package task

import (
	"encoding/json"
	"fmt"
	"time"

	"example/status"
)

//
func ExampleTaskMarshalJSON() {
	d, _ := time.Parse(time.RFC3339, "2020-11-30T10:09:27.2389529Z")
	t := Task{
		Title:    "Laundry",
		Status:   status.Done,
		Deadline: &Deadline{d},
		SubTasks: []Task{},
	}
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"Done","deadline":1606730967}
}

//
func ExampleTaskUnmarshalJSON() {
	b := []byte(`{"title":"Laundry","status":"Done","deadline":1606730967}`)
	var t Task
	json.Unmarshal(b, &t)
	fmt.Println(t)
	// Output:
	// {Laundry 3 2020-11-30 10:09:27 +0000 UTC []}
}
