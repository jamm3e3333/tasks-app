package repository

import (
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
)

type TaskRepositoryPort interface {
	TaskById(id uuid.UUID) (<-chan entity.TaskEntity, <-chan error)
	Tasks() (<-chan entity.TaskEntity, <-chan bool)
	Create(task entity.TaskEntity) error
	Update(task entity.TaskEntity) error
	DeleteById(id uuid.UUID) error
}
