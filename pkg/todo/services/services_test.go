package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"github.com/xfyuan/go-yesteaser/pkg/gspec"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
	"github.com/xfyuan/go-yesteaser/pkg/todo/daos"
)

var _ = Describe("Todo Services", func() {
	var (
		service todo.Service
		data *models.Todo
		err error
	)

	BeforeEach(func() {
		app.DB = gspec.ResetDB()
		service = NewTodoService(daos.NewTodoDao(app.DB))
	})

	Describe("with todo records", func() {
		Context("when exists", func() {
			BeforeEach(func() {
				app.DB.Create(&models.Todo{
					Title:       "Golang",
					Description: "Google's language",
				})
				data, err = service.Get(1)
			})

			It("should has no error", func() {
				Expect(err).To(BeNil())
			})

			It("should has correct record", func() {
				Expect(data.Title).To(Equal("Golang"))
				Expect(data.Description).To(Equal("Google's language"))
			})
		})

		Context("when not exists", func() {
			BeforeEach(func() {
				data, err = service.Get(9999)
			})

			It("should has error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should has empty record", func() {
				Expect(data.Title).To(BeEmpty())
				Expect(data.Description).To(BeEmpty())
			})
		})
	})
})
