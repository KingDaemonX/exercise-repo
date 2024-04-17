package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Task represents a todo item
type Task struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Done  bool      `json:"done"`
}

type TaskBody struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var localMemoryStore map[uuid.UUID]*Task = make(map[uuid.UUID]*Task)

func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]*Task, 0, len(localMemoryStore))
	for _, task := range localMemoryStore {
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func handleGetTask(w http.ResponseWriter, id uuid.UUID) {
	task, ok := localMemoryStore[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var TaskBody TaskBody

	err := json.NewDecoder(r.Body).Decode(&TaskBody)
	if err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	newTask := &Task{
		ID:    uuid.New(),
		Title: TaskBody.Title,
		Done:  TaskBody.Done,
	}

	localMemoryStore[newTask.ID] = newTask
	w.WriteHeader(http.StatusCreated)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	task, ok := localMemoryStore[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Title = updatedTask.Title
	task.Done = updatedTask.Done
	w.WriteHeader(http.StatusOK)
}

func handleDeleteTask(w http.ResponseWriter, id uuid.UUID) {
	_, ok := localMemoryStore[id]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	delete(localMemoryStore, id)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/tasks", handleGetTasks)
	http.HandleFunc("/tasks/create", handleCreateTask)
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.URL.Path[len("/tasks/"):])
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handleGetTask(w, id)
		case http.MethodPut:
			handleUpdateTask(w, r, id)
		case http.MethodDelete:
			handleDeleteTask(w, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
