package cmd

import (
	"github.com/spf13/cobra"
)

var password string
var key string
var value string
var yamlfile string
var envfile string
var templatefile string
var outfile string

var RootCmd = &cobra.Command{
	Use:   "gotmpl",
	Short: "gotmpl is a template tool that supports encrypted data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func Execute() {
	RootCmd.Execute()
}
