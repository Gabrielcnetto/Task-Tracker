package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *Task) CreatedTime() {
	t.CreatedAt = time.Now()
}
func (t *Task) UpdatedTime() {
	t.UpdatedAt = time.Now()
}
func (t *Task) GenerateId() {
	myuuid := uuid.New().String()
	t.Id = myuuid
}
