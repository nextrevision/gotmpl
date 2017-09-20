package cmd

import (
	"fmt"

	"github.com/nextrevision/gotmpl/pkg"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "takes a plain text value and encrypts it",
	Long: `Takes a plain text value and encrypts it (optionally inserting it into a YAML file).

  Examples:

  // encrypts 'test' to stdout
  gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v test
  // encrypts 'test' to stdout to be pasted into a YAML file
  gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v test -k mykey
  // encrypts 'test' and writes it to 'mykey' in a YAML file
  gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v test -k mykey -y myyamlfile.yml

	`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(password) != 32 {
			return fmt.Errorf("Password must be exactly 32 characters, not %d. Try running 'gotmpl genpasswd'", len(password))
		}
		if value == "" {
			return fmt.Errorf("Must supply a value to encrypt")
		}
		if yamlfile != "" && envfile != "" {
			return fmt.Errorf("Vars file can only either be YAML or ENV")
		}
		if yamlfile != "" {
			if !pkg.FileExists(yamlfile) {
				return fmt.Errorf("No such vars file: %s", yamlfile)
			}
			if key == "" {
				return fmt.Errorf("Must supply key when inserting into a vars file")
			}
		}
		if envfile != "" {
			if !pkg.FileExists(envfile) {
				return fmt.Errorf("No such vars file: %s", envfile)
			}
			if key == "" {
				return fmt.Errorf("Must supply key when inserting into a vars file")
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		encryptedValue, err := pkg.EncryptString(value, password)
		if err != nil {
			return err
		}

		if yamlfile != "" {
			vars, err := pkg.LoadYAMLFile(yamlfile)
			if err != nil {
				return err
			}

			vars[key] = fmt.Sprintf("ENC|%s", encryptedValue)
			if err = pkg.WriteYAMLFile(vars, yamlfile); err != nil {
				return err
			}
			fmt.Printf("Variable %s inserted into %s\n", key, yamlfile)
		} else if envfile != "" {
			vars, err := pkg.LoadEnvFile(envfile)
			if err != nil {
				return err
			}

			vars[key] = fmt.Sprintf("ENC|%s", encryptedValue)
			if err = pkg.WriteEnvFile(vars, envfile); err != nil {
				return err
			}
			fmt.Printf("Variable %s inserted into %s\n", key, envfile)
		} else {
			if key == "" {
				fmt.Printf(encryptedValue)
			} else {
				fmt.Printf("%s: ENC|%s", key, encryptedValue)
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&password, "password", "p", "", "password to encrypt with")
	encryptCmd.Flags().StringVarP(&value, "value", "v", "", "value to encrypt")
	encryptCmd.Flags().StringVarP(&key, "key", "k", "", "key to assign value to")
	encryptCmd.Flags().StringVarP(&yamlfile, "yaml", "y", "", "yaml file to insert result into (optional)")
	encryptCmd.Flags().StringVarP(&envfile, "env", "e", "", "env vars file to use with template (optional)")
}
