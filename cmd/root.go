package cmd

import (
	"github.com/localghost/healthy/checker"
	"github.com/localghost/healthy/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "healthy",
	Short: "Light health checker with HTTP REST API",
	Run: func(cmd *cobra.Command, args []string) {
		checker := checker.New(viper.Get("checks"))
		checker.Start()

		server.New(checker).Start()
	},
}

var cfgFile string

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(readConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./healthy.yml)")
	rootCmd.Flags().String("address", "127.0.0.1", "address to listen on")
	rootCmd.Flags().Int("port",  8199, "port to listen on")
	viper.BindPFlag("server.port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.address", rootCmd.Flags().Lookup("address"))
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
