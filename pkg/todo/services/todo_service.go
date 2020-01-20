package services

import (
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
)

type TodoService struct {
	dao todo.Dao
}

func NewTodoService(d todo.Dao) todo.Service {
	return &TodoService{
		dao: d,
	}
}

func (s *TodoService)Get(id int64) (*models.Todo, error) {
	return s.dao.Get(id)
}
