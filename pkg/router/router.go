package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_todoRoutes "github.com/xfyuan/go-yesteaser/pkg/todo/routes"
)

func Initialize(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	_todoRoutes.InitRoutes(db, r)

	return r
}