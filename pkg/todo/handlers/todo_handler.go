package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TodoHandler struct {
	service todo.Service
}

func NewTodoHandler(r *gin.Engine, s todo.Service) {
	handler := &TodoHandler{
		service: s,
	}

	v := r.Group("/api/v1")
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
