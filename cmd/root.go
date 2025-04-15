/*
Copyright © 2025 Naufal Khairil Imami
*/
package cmd

import (
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filewatcher",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working dir, %s", err)
	}

	configFile := path.Join(wd, "filewatcher.yaml")
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	logLevel := log.InfoLevel
	logLevelCfg := viper.GetString("log.level")
	if logLevelCfg != "" {
		parseLevel, err := log.ParseLevel(logLevelCfg)
		if err != nil {
			log.Fatalf("Failed to parse log level, %s", err)
		}

		logLevel = parseLevel
	}
	log.SetLevel(logLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})

}
