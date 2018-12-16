package cmd

import (
	"github.com/localghost/healthy/checker"
	"github.com/localghost/healthy/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use: "healthy",
	Short: "Light health checker with HTTP REST API",
	Run: func(cmd *cobra.Command, args []string) {
		checker, err := checker.NewChecker(viper.Get("checks"))
		if err != nil {
			panic(err)
		}
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
	rootCmd.Flags().String("listen_on", "127.0.0.1:8199", "address to listen on")
	viper.BindPFlag("server.listen_on", rootCmd.Flags().Lookup("listen_on"))
	viper.SetDefault("checker.interval", "10s")
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
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
