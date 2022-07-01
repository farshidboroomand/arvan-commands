package cmd

import (
	"arvan_internal_cli/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var DomainName string
var DomainId string
var Config = config.GetConfigInfo()

var helpDescriptions = map[string]string{
	"pageRul-list":   "Get All domain pageRule",
	"pageRul-delete": "Delete all pageRules ",
}

var rootCmd = &cobra.Command{
	Use:   "ar-cli",
	Short: "This package provides a unified command line interface to Arvan CDN Services.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
