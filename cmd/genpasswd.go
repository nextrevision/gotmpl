package cmd

import (
	"fmt"

	"github.com/nextrevision/gotmpl/pkg"
	"github.com/spf13/cobra"
)

var length int

var genpasswdCmd = &cobra.Command{
	Use:   "genpasswd",
	Short: "generates a compliant 32 character password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.NewPassword(length))
	},
}

func init() {
	RootCmd.AddCommand(genpasswdCmd)
	genpasswdCmd.Flags().IntVarP(&length, "length", "l", 32, "length of password")
}
