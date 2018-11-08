package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "healthy",
	Short: "Light health checker with HTTP REST API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Doin stuff")
	},
}
var cfgFile string

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(readConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./healthy.yml)")
}

func readConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("healthy")
		viper.AddConfigPath("/etc/healthy")
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix("HEALTHY")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}