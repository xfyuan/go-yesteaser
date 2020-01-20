package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cobra"
	"github.com/xfyuan/go-yesteaser/pkg/app"
	"github.com/xfyuan/go-yesteaser/pkg/models"
	"github.com/xfyuan/go-yesteaser/pkg/router"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting http server ...")

		dsn := app.GenerateDSN()

		app.DB, app.DBErr = gorm.Open("postgres", dsn)
		if app.DBErr != nil {
			panic(app.DBErr)
		}

		defer func() {
			err := app.DB.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		app.DB.AutoMigrate(&models.Todo{})
		log.Println("Successfully connected to database")

		r := router.Initialize(app.DB)

		if err := r.Run(fmt.Sprintf(":%v", "1234")); err != nil {
			panic(fmt.Errorf("gin run failed: [%s]", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

