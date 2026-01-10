package services

import (
	"fmt"
	"tasks-cli/models"
)

func Update(TaskId string, status int8) error {
	tasks, err := GetTasks()
	if err != nil {
		return err
	}
	updatedList := []models.Task{}
	for _, item := range tasks {
		if item.Id == TaskId {

			item.Status = status
			item.UpdatedTime()
			fmt.Println(item)
		}
		updatedList = append(updatedList, item)
	}
	newSave := SaveTask(updatedList)
	if newSave != nil {
		return newSave
	}
	return nil
}
