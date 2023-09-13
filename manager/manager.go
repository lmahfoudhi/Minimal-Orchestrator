package manager

import (
	"fmt"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]Task
	EventDb       map[string][]TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
