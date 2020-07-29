package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"github.com/xfyuan/go-yesteaser/pkg/gspec"
	"github.com/xfyuan/go-yesteaser/pkg/middlewares"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/todo/daos"
	"github.com/xfyuan/go-yesteaser/pkg/todo/services"
)

var _ = Describe("Todo Handlers", func() {
	var (
		r *gin.Engine
		h TodoHandler
	)

	BeforeEach(func() {
		app.DB = gspec.ResetDB()

		r = gspec.NewRouter()
		r.Use(gspec.SetAuthHeader())
		r.Use(middlewares.Auth())

		h = TodoHandler{
			service: services.NewTodoService(daos.NewTodoDao(app.DB)),
		}
	})

	Describe("#Show", func() {
		Context("with record", func() {
			BeforeEach(func() {
				app.DB.Create(&models.Todo{
					Title:       "Golang",
					Description: "Google's language",
				})
			})

			It("should get a todo successfully", func() {
				r.Handle("GET", "/todos/:id", h.Show)
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/todos/1", nil)
				r.ServeHTTP(res, req)

				body, _ := ioutil.ReadAll(res.Body)
				var expect map[string]interface{}
				if err := json.Unmarshal(body, &expect); err != nil {
					panic(err)
				}

				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(expect["title"]).To(Equal("Golang"))
				Expect(expect["description"]).To(Equal("Google's language"))

			})
		})

		Context("without record", func() {
			It("should not found a todo", func() {
				r.Handle("GET", "/todos/:id", h.Show)
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/todos/1", nil)
				r.ServeHTTP(res, req)

				body, _ := ioutil.ReadAll(res.Body)

				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(body).To(BeEmpty())
			})
		})
	})
})
