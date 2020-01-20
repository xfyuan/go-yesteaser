package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_todoHandlers "github.com/xfyuan/go-yesteaser/pkg/todo/handlers"
)

func Initialize(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	_todoHandlers.Initialize(db, r)

	return r
}