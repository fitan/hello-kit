/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"hello/cmd/app"

	"github.com/spf13/cobra"
)

var (
	gitHash   string
	gitTag    string
	buildTime string
	goVersion string
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @query.collection.format multi
// @securityDefinitions.basic  BasicAuth

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use: "app",
	//	Short: "A brief description of your command",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		if version {
	//			fmt.Printf("Git Tag: %s \n", gitTag)
	//			fmt.Printf("Git Commit hash: %s \n", gitHash)
	//			fmt.Printf("Build TimeStamp: %s \n", buildTime)
	//			fmt.Printf("GoLang Version: %s \n", goVersion)
	//			return
	//		}
	//		app.RunApp(confName)
	//	},
}

var appStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start app",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunApp(confName)
	},
}

var appPathPutInStorageCmd = &cobra.Command{
	Use:     "pathPutInStorage",
	Aliases: []string{"ppis"},
	Short:   "PathPutInStorage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.PathPutInStorage(confName)
	},
}

var appCasbinCmd = &cobra.Command{
	Use:     "casbin input",
	Aliases: []string{"casbin"},
	Short:   "casbin input",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.CasbinInput(confName)
	},
}

var appVersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "echo app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Git Tag: %s \n", gitTag)
		fmt.Printf("Git Commit hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
	},
}

var confName string

func init() {
	appCmd.AddCommand(appStartCmd, appVersionCmd, appPathPutInStorageCmd, appCasbinCmd)
	rootCmd.AddCommand(appCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	appPathPutInStorageCmd.Flags().StringVarP(&confName, "confName", "c", "dev", "config name")
	appStartCmd.Flags().StringVarP(&confName, "confName", "c", "dev", "config name")
	appCasbinCmd.Flags().StringVarP(&confName, "confName", "c", "dev", "config name")
}
