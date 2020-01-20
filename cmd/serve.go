package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	_todoApi "github.com/xfyuan/go-yesteaser/pkg/todo/api"
	_todoDao "github.com/xfyuan/go-yesteaser/pkg/todo/dao"
	_todoService "github.com/xfyuan/go-yesteaser/pkg/todo/service"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting http server ...")

		dsn := generateDSN()

		db, dbErr := gorm.Open("postgres", dsn)
		if dbErr != nil {
			panic(dbErr)
		}

		defer func() {
			err := db.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		db.AutoMigrate(&models.Todo{})
		log.Println("Successfully connected to database")

		r := gin.New()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())

		td := _todoDao.NewTodoDao(db)
		ts := _todoService.NewTodoService(td)
		_todoApi.NewTodoHandler(r, ts)

		if err := r.Run(fmt.Sprintf(":%v", "1234")); err != nil {
			panic(fmt.Errorf("gin run failed: [%s]", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func generateDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)
}

