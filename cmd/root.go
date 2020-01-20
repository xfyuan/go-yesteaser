package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/xfyuan/go-yesteaser/pkg/app"
  "os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "go-yesteaser",
  Short: "Yesteaser is a Go project layout boilerplate",
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {
    _ = cmd.Usage()
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  app.LoadConfig("")
}

