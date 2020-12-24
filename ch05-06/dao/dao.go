package dao

import (
	"fmt"

	"example/task"
)

// ID is identity of task.Task object.
type ID string

// DAO is interface for Data Access Object.
type DAO interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}

// MemoryDAO is implementation of DAO with in-memory methods.
type MemoryDAO struct {
	tasks  map[ID]task.Task
	nextID int64
}

// NewMemoryDAO returns new instance of MemoryDAO, as pointer.
func NewMemoryDAO() *MemoryDAO {
	return &MemoryDAO{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}
}

// Get returns task.Task of corresponding id, error is nil if object exists.
func (m *MemoryDAO) Get(id ID) (task.Task, error) {
	t, ok := m.tasks[id]
	if !ok {
		return task.Task{}, fmt.Errorf("MemoryDAO.Get: no task found for %v", id)
	}
	return t, nil
}

// Put updates existing object.
func (m *MemoryDAO) Put(id ID, t task.Task) error {
	if _, ok := m.tasks[id]; !ok {
		return fmt.Errorf("MemoryDAO.Put: no task found for %v", id)
	}
	m.tasks[id] = t
	return nil
}

// Post add new object and returns its id assigned.
func (m *MemoryDAO) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprintf("%d", m.nextID))
	m.tasks[id] = t
	m.nextID++
	return id, nil
}

// Delete deletes existing object
func (m *MemoryDAO) Delete(id ID) error {
	if _, ok := m.tasks[id]; !ok {
		return fmt.Errorf("MemoryDAO.Delete: no task found for %v", id)
	}
	delete(m.tasks, id)
	return nil
}
