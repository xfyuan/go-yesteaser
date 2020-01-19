package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "os"
)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "go-yesteaser",
  Short: "Yesteaser is a Go project layout boilerplate",
  // Uncomment the following line if your bare application
  // has an action associated with it:
  	Run: func(cmd *cobra.Command, args []string) {
  	  cmd.Usage()
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

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/config.yml)")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    viper.AddConfigPath("config")
    viper.SetConfigName("config")
  }

  viper.SetEnvPrefix("yestea")
  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err != nil {
    panic(fmt.Errorf("using config failed: [%s]", err))
  }
  fmt.Println("Using config file:", viper.ConfigFileUsed())

  env := viper.GetString("ENV")
  if env == "" {
    viper.SetConfigName("dev")
  }
  viper.SetConfigName(env)
  if err := viper.MergeInConfig(); err != nil {
    panic(fmt.Errorf("merge environment config failed: [%s]", err))
  }
  fmt.Println("Using config file:", viper.ConfigFileUsed())
}

