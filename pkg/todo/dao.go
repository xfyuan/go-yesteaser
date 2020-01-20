package todo

import (
	"github.com/xfyuan/go-yesteaser/pkg/models"
)

type Dao interface {
	Get(id int64) (*models.Todo, error)
}