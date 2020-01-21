package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xfyuan/go-yesteaser/pkg/middlewares"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
	_todoDaos "github.com/xfyuan/go-yesteaser/pkg/todo/daos"
	_todoServices "github.com/xfyuan/go-yesteaser/pkg/todo/services"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TodoHandler struct {
	service todo.Service
}

func Initialize(db *gorm.DB, r *gin.Engine)  {
	s := _todoServices.NewTodoService(_todoDaos.NewTodoDao(db))
	NewTodoHandler(r, s)
}

func NewTodoHandler(r *gin.Engine, s todo.Service) {
	handler := &TodoHandler{
		service: s,
	}

	v := r.Group("/api/v1").Use(middlewares.Auth())
	{
		v.GET("/todos/:id", handler.Show)
	}
}

func (h *TodoHandler) Show(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	if data, err := h.service.Get(int64(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
