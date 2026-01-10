package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"tasks-cli/models"
)

const repo = "./repositories/data.json"

func verifyDir() int {

	dir := filepath.Dir(repo)
	//aqui garante que a pasta existe
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return 0
		}
	}
	//agora verifica se o arquivo existe
	if _, err := os.Stat(repo); os.IsNotExist(err) {
		file, err := os.Create(repo)
		if err != nil {
			return 1
		}
		defer file.Close()
		return 2

	} else {
		return 3
	}

}

func AddTask(task models.Task) error {
	dirStatus := verifyDir()
	switch dirStatus {
	case 2:
		return SaveTask([]models.Task{task})
	case 3:
		err := putTask(task)
		return err
	default:
		return fmt.Errorf("Ocorreu um erro")
	}

}
func SaveTask(task []models.Task) error {
	data, err := json.MarshalIndent(task, "", "")
	if err != nil {
		return err
	}
	err = os.WriteFile(repo, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func putTask(task models.Task) error {
	tasks, err := GetTasks()
	if err != nil {
		return err
	}
	err = SaveTask(append(tasks, task))
	if err != nil {
		return err
	}
	return nil
}

func GetTasks() ([]models.Task, error) {
	data, err := os.ReadFile(repo)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
