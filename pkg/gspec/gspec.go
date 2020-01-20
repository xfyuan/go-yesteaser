package gspec

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	if err := os.Setenv("YESTEA_ENV", "test"); err != nil {
		panic(fmt.Errorf("set test env failed: [%s]", err))
	}

	app.LoadConfig(RootPath() + "config")

	dsn := app.GenerateDSN()

	app.DB, app.DBErr = gorm.Open("postgres", dsn)
	if app.DBErr != nil {
		panic(app.DBErr)
	}

	app.DB.AutoMigrate(&models.Todo{})
}

func RootPath() string {
	// find current path
	_, filename, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(filename)
	rootDir := currentPath + "/../../"
	return rootDir
}

func ResetDB() *gorm.DB {
	app.DB.DropTableIfExists(&models.Todo{})
	app.DB.AutoMigrate(&models.Todo{})
	return app.DB
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func SetAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Authorization", viper.GetString("apikey"))
	}
}

