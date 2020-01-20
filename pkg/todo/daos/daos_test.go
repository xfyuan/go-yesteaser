package daos

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"github.com/xfyuan/go-yesteaser/pkg/gspec"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/todo"
)

var _ = Describe("Daos", func() {
	var (
		dao todo.Dao
		data *models.Todo
		err error
	)

	BeforeEach(func() {
		app.DB = gspec.ResetDB()
	})

	Describe("with todo records", func() {
		Context("when exists", func() {
			BeforeEach(func() {
				app.DB.Create(&models.Todo{
					Title:       "Golang",
					Description: "Google's language",
				})
				dao = NewTodoDao(app.DB)
				data, err = dao.Get(1)
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
				dao = NewTodoDao(app.DB)
				data, err = dao.Get(9999)
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
