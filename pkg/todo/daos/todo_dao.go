package daos

import (
	"github.com/jinzhu/gorm"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
)

type TodoDao struct {
	DB *gorm.DB
}

func NewTodoDao(db *gorm.DB) todo.Dao {
	return &TodoDao{db}
}

func (d *TodoDao)Get(id int64) (*models.Todo, error) {
	var data models.Todo

	err := d.DB.Where("id = ?", id).
		First(&data).
		Error

	return &data, err
}
