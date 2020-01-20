package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_todoApis "github.com/xfyuan/go-yesteaser/pkg/todo/apis"
	_todoDaos "github.com/xfyuan/go-yesteaser/pkg/todo/daos"
	_todoServices "github.com/xfyuan/go-yesteaser/pkg/todo/services"
)

func InitRoutes(db *gorm.DB, r *gin.Engine)  {
	s := _todoServices.NewTodoService(_todoDaos.NewTodoDao(db))
	_todoApis.NewTodoHandler(r, s)
}
