package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/err0r500/go-realworld-clean/implem/gin.server"
	"github.com/err0r500/go-realworld-clean/implem/jwt.authHandler"
"github.com/err0r500/go-realworld-clean/implem/logrus.logger"
	"github.com/err0r500/go-realworld-clean/infra"
	"github.com/err0r500/go-realworld-clean/uc"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "go-realworld-clean",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(infra.CobraInitialization)

	infra.LoggerConfig(rootCmd)
	infra.ServerConfig(rootCmd)
	infra.DatabaseConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)

	authHandler := jwt.NewTokenHandler(viper.GetString("jwt.Salt"))
	
	server.NewRouter(
		uc.NewHandler(
			logger.NewLogger("TEST",
				viper.GetString("log.level"),
				viper.GetString("log.format"),
			),
			nil, //fixme : not implemented yet
			nil, //fixme : not implemented yet
			authHandler,
		),
		authHandler,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
