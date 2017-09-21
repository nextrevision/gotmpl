package cmd

import (
	"fmt"

	"github.com/nextrevision/gotmpl/pkg"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypts vars or files",
	Long: `Takes input as either a value or a file to decrypt.

  Examples:

  // decrypts an entire file to vars.yml.unenc
  gotmpl decrypt -p password -y vars.yml
  // decrypts a single value to stdout
  gotmpl decrypt -p password -v NjExMDlmY2QyOTc4MTg0MTFkZDBhYjM5OmEzMTFkMjU4ODU5M2U3ZWJmZTdjNDMxMWEzY2VhNzlmNTJiOWUzZGI=

	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if password == "" {
			return fmt.Errorf("Must supply a password")
		}
		if yamlfile != "" && envfile != "" {
			return fmt.Errorf("Vars file can only either be YAML or ENV")
		}
		if value == "" && yamlfile == "" && envfile == "" {
			return fmt.Errorf("Must supply either a value or vars file to decrypt")
		}
		if yamlfile != "" && !pkg.FileExists(yamlfile) {
			return fmt.Errorf("No such vars file: %s", yamlfile)
		}
		if envfile != "" && !pkg.FileExists(envfile) {
			return fmt.Errorf("No such vars file: %s", envfile)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if value != "" {
			decryptedString, err := pkg.DecryptString(value, password)
			if err != nil {
				return err
			}
			fmt.Printf("%s", decryptedString)
		} else {
			if yamlfile != "" {
				vars, err := pkg.LoadYAMLFile(yamlfile)
				if err != nil {
					return err
				}
				decryptedValues, err := pkg.DecryptValues(vars, password)
				if err != nil {
					return err
				}
				if err = pkg.WriteYAMLFile(decryptedValues, fmt.Sprintf("%s.unenc", yamlfile)); err != nil {
					return err
				}
				fmt.Printf("File decrypted to %s.unenc", yamlfile)
			} else if envfile != "" {
				vars, err := pkg.LoadEnvFile(envfile)
				if err != nil {
					return err
				}
				decryptedValues, err := pkg.DecryptValues(vars, password)
				if err != nil {
					return err
				}
				if err = pkg.WriteEnvFile(decryptedValues, fmt.Sprintf("%s.unenc", envfile)); err != nil {
					return err
				}
				fmt.Printf("File decrypted to %s.unenc", envfile)
			}

		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&password, "password", "p", "", "password to decrypt vars with")
	decryptCmd.Flags().StringVarP(&value, "value", "v", "", "encrypted value to decrypt")
	decryptCmd.Flags().StringVarP(&yamlfile, "yaml", "y", "", "yaml vars file to decrypt")
	decryptCmd.Flags().StringVarP(&envfile, "env", "e", "", "env vars file to decrypt")
}
