package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
	"github.com/jamm3e3333/tasks-app/pkg/logger"
)

type TaskRepository struct {
	memory *map[uuid.UUID]entity.TaskEntity
	l      logger.Logger
}

func NewTaskRepository(l logger.Logger) *TaskRepository {
	newRepoMemory := map[uuid.UUID]entity.TaskEntity{}

	return &TaskRepository{
		&newRepoMemory,
		l,
	}
}

func (tr *TaskRepository) TaskById(id uuid.UUID) (<-chan entity.TaskEntity, <-chan error) {
	ch := make(chan entity.TaskEntity)
	ech := make(chan error)

	go func() {
		defer close(ch)
		defer close(ech)

		task, isExist := (*tr.memory)[id]

		if !isExist {
			err := errors.New(fmt.Sprintf("task with id %s doesn't exist", id))
			tr.l.Debug(err)
			ech <- err
		} else {
			ch <- task
		}
	}()

	return ch, ech
}

func (tr *TaskRepository) Tasks() (<-chan entity.TaskEntity, <-chan bool) {
	ch := make(chan entity.TaskEntity)
	done := make(chan bool)

	go func() {
		defer close(ch)
		defer close(done)

		if tr.isEmpty() {
			done <- true
			return
		}

		for _, value := range *tr.memory {
			ch <- value
		}
		done <- true
	}()

	return ch, done
}

func (tr *TaskRepository) isEmpty() bool {
	return len(*tr.memory) == 0
}

func (tr *TaskRepository) notFoundError(id uuid.UUID) error {
	return errors.New(fmt.Sprintf("task with id: %s doesn't exist", id))
}

func (tr *TaskRepository) Create(task entity.TaskEntity) error {
	(*tr.memory)[task.Id] = task
	return nil
}

func (tr *TaskRepository) Update(task entity.TaskEntity) error {
	taskToBeUpdated, isExist := (*tr.memory)[task.Id]

	if !isExist {
		return tr.notFoundError(taskToBeUpdated.Id)
	}
	(*tr.memory)[taskToBeUpdated.Id] = task

	return nil
}

func (tr *TaskRepository) DeleteById(id uuid.UUID) error {

	if isEmpty := tr.isEmpty(); isEmpty {
		return tr.notFoundError(id)
	}

	task, err := tr.TaskById(id)
	select {
	case <-err:
		return tr.notFoundError(id)
	case <-task:
		delete(*tr.memory, id)
		return nil
	}
}
