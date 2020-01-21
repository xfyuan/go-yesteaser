package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			app.NewError(c, http.StatusUnauthorized, errors.New("Authorization header is required!"))
			c.Abort()
		}
		if authHeader != viper.Get("apikey") {
			app.NewError(c, http.StatusUnauthorized, fmt.Errorf("Not authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}

