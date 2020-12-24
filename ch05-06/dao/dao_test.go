package dao

import (
	"example/status"
	"example/task"
	"reflect"
	"testing"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMemoryDAO(t *testing.T) {
	defer func() {
		if c := recover(); c != nil {
			t.Errorf("unexpected panic: %+v", c)
		}
	}()

	m := NewMemoryDAO()
	t1 := task.Task{
		Title:  "Laundry",
		Status: status.WIP,
	}
	id, err := m.Post(t1)
	panicErr(err)

	t2, err := m.Get(id)
	panicErr(err)
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("expecting t1 == t2 but found different")
	}

	t3 := task.Task{
		Title:  "Baking",
		Status: status.Done,
	}
	err = m.Put(id, t3)
	panicErr(err)
	t4, err := m.Get(id)
	panicErr(err)
	if !reflect.DeepEqual(t3, t4) {
		t.Errorf("expecting t3 == t4 but found different")
	}

	err = m.Delete(id)
	panicErr(err)
}
