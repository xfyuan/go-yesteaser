package middlewares

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xfyuan/go-yesteaser/pkg/gspec"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Middlewares", func() {
	var (
		r *gin.Engine
	)

	BeforeEach(func() {
		r = gspec.NewRouter()
	})

	Describe("Auth using api key", func() {
		Context("without authorization header", func() {
			It("should not pass", func() {
				r.GET("/", Auth())
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(res, req)

				Expect(res.Code).To(Equal(http.StatusUnauthorized))
			})
		})

		Context("with api key in authorization header", func() {
			It("should pass", func() {
				r.GET("/", gspec.SetAuthHeader(), Auth())
				res := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(res, req)

				Expect(res.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
