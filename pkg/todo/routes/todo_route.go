package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_todoDaos "github.com/xfyuan/go-yesteaser/pkg/todo/daos"
	_todoHandlers "github.com/xfyuan/go-yesteaser/pkg/todo/handlers"
	_todoServices "github.com/xfyuan/go-yesteaser/pkg/todo/services"
)

func InitRoutes(db *gorm.DB, r *gin.Engine)  {
	s := _todoServices.NewTodoService(_todoDaos.NewTodoDao(db))
	_todoHandlers.NewTodoHandler(r, s)
}
