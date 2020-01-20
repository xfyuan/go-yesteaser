package api

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
	ownService todo.Service
}

func NewTodoHandler(r *gin.Engine, s todo.Service) {
	handler := &TodoHandler{
		ownService: s,
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos/:id", handler.Get)
	}
}

func (h *TodoHandler)Get(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	if data, err := h.ownService.Get(int64(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
