package commands

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var tasks = readTaskFromJSONFile()

func Add(lask string) {
	lastTaskId := 0
	if len(tasks) > 0 {
		lastTaskId = tasks[len(tasks)-1].ID
	}
	newtask := &Task{
		ID:          lastTaskId + 1,
		Description: lask,
		Status:      "TODO",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"), // установка времени
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, newtask)
	writeTaskToJSONFile(tasks)
	fmt.Println("Task added successfully, ID:", newtask.ID)
}

func findById(id string) (*Task, int) {
	num, err := strconv.Atoi(id) // перевод типов из строки в число
	if err != nil {
		fmt.Println("Error, u need to write second argument 'ID'")
	}
	for index, task := range tasks {
		if task.ID == num {
			return task, index
		}
	}
	return nil, -1
}

func MarkInProgress(id string) {
	idw, index := findById(id)
	if idw == nil {
		fmt.Println("Task not found")
	}
	idw.Status = "In Progress"
	tasks[index] = idw
	writeTaskToJSONFile(tasks)
	fmt.Println("Mark is changed")
}
func MarkDone(id string) {
	idw, index := findById(id)
	if idw == nil {
		fmt.Println("Task not found")
	}
	idw.Status = "DONE"
	tasks[index] = idw
	writeTaskToJSONFile(tasks)
	fmt.Println("Mark is changed")
}

func Update(id string, lask string) {
	idw, index := findById(id)
	if idw == nil {
		fmt.Println("Task not found")
	}
	idw.Description = lask
	idw.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	tasks[index] = idw
	writeTaskToJSONFile(tasks)
	fmt.Println("Task is changed correctly")
}

func Delete(id string) {
	idw, index := findById(id)
	if idw == nil {
		fmt.Println("Task not found", index)
	}
	for i, v := range tasks {
		if v == idw {
			copy(tasks[i:], tasks[i+1:])
			tasks[len(tasks)-1] = nil
			tasks = tasks[:len(tasks)-1]
		}

	}
	writeTaskToJSONFile(tasks)
	fmt.Println("Delete completed")
}

func List() {
	fmt.Printf("%-5s %-50s %-15s\n", "ID", "Description", "Status")
	for _, task := range tasks {
		fmt.Printf("%-5d %-50s %-15s\n", task.ID, task.Description, task.Status)
	}
}
