package services

import (
	"tasks-cli/models"
)

func DeletTasks(taskId string) error {
	tasks, err := GetTasks()
	if err != nil {
		return err
	}
	newList := []models.Task{}
	for _, item := range tasks {
		if item.Id != taskId {
			newList = append(newList, item)
			continue
		}
		continue
	}

	newSave := SaveTask(newList)
	if newSave != nil {
		return newSave
	}
	return nil
}
