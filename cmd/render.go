package cmd

import (
	"fmt"

	"github.com/nextrevision/gotmpl/pkg"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "renders a template file",
	Long: `Renders a template file with inputs from the environment or a vars file.

IMPORTANT: environment variables always overwrite values from the vars file.

  Examples:

  // render 'template.tmpl' using only environment variables to stdout
  gotmpl render -t template.tmpl
  // render using env vars and keys from 'vars.yml'
  gotmpl render -t template.tmpl -y vars.yml
  // decrypt any variables (inline or in a file) using a password
  gotmpl render -t template.tmpl -y vars.yml -p password
  // write the rendered template to a specific file
  gotmpl render -t template.tmpl -o template.txt

	`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if !pkg.FileExists(templatefile) {
			return fmt.Errorf("Must supply a template file")
		}
		if yamlfile != "" && envfile != "" {
			return fmt.Errorf("Vars file can only either be YAML or ENV")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		vars := make(map[string]interface{})
		var err error
		if yamlfile != "" {
			vars, err = pkg.LoadYAMLFile(yamlfile)
			if err != nil {
				return err
			}
		}
		if envfile != "" {
			vars, err = pkg.LoadEnvFile(envfile)
			if err != nil {
				return err
			}
		}

		return pkg.ConvertTemplate(templatefile, vars, outfile, password)
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringVarP(&templatefile, "template", "t", "", "template to render")
	renderCmd.Flags().StringVarP(&outfile, "out", "o", "", "output file (default: stdout)")
	renderCmd.Flags().StringVarP(&password, "password", "p", "", "password decrypt secret values with (optional)")
	renderCmd.Flags().StringVarP(&yamlfile, "yaml", "y", "", "yaml vars file to use with template (optional)")
	renderCmd.Flags().StringVarP(&envfile, "env", "e", "", "env vars file to use with template (optional)")
}
