package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath() string {
	fle, error := os.Getwd() // определяем текущую директорию

	if error != nil { // обработка ошибок
		fmt.Println("Error getting filePath")
	}
	return filepath.Join(fle, "tasks.json") // возвращает обьединение пути и названия json-файла
}

func readTaskFromJSONFile() []*Task {
	tasks := []*Task{}                // инициализация пустого списка задач
	taskFilePath := getFilePath()     // получаем путь к файлу
	_, error := os.Stat(taskFilePath) // проверка существования файла
	if error != nil {
		if os.IsNotExist(error) { // если не сущетсвует файл
			file, error := os.Create(taskFilePath)                         // создает файл
			os.WriteFile(taskFilePath, []byte("[]"), os.ModeAppend.Perm()) // пишет пустой json массив
			if error != nil {
				fmt.Println("Error creating tasks file: ", error)
			}

			file.Close()
		} else {
			fmt.Println("Error checking if file exists: ", error)
		}
	}

	bytes, error := os.ReadFile(taskFilePath) // читает весь файл

	if error != nil {
		fmt.Println("Error reading tasks from file: ", error)
	}
	error = json.Unmarshal(bytes, &tasks) // парсит json (преобразует json данные в структуры Go (тип Task))

	if error != nil {
		fmt.Println("Error unmarshalling tasks: ", error)
	}

	return tasks
}

func writeTaskToJSONFile(tasks []*Task) {
	jsonData, err := json.Marshal(tasks) // конвертирует задачи в json
	taskFilePath := getFilePath()

	if err != nil {
		fmt.Println("Error marshalling tasks: ", err)
	}

	err = os.WriteFile(taskFilePath, jsonData, os.ModeAppend.Perm()) // записываем данные в файл

	if err != nil {
		fmt.Println("Error writing tasks to file", err)
	}
}
