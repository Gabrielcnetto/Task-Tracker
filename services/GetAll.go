package services

import "tasks-cli/models"

func GetAllTasks(filter int) ([]models.Task, error) {
	tasks, err := GetTasks()
	if filter == 0 {
		if err != nil {
			return nil, err
		}
		return tasks, nil
	}
	filteredList := GetTasksWithFilter(filter, tasks)
	return filteredList, nil
}

func GetTasksWithFilter(filter int, tasks []models.Task) []models.Task {
	filteredList := []models.Task{}
	for _, item := range tasks {
		if item.Status == int8(filter) {
			filteredList = append(filteredList, item)
		}
	}
	return filteredList
}
