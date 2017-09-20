package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// ConvertTemplate renders a template
func ConvertTemplate(templatefile string, vars map[string]interface{}, outfile string, password string) error {
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"ENC": func(encryptedString string) string {
			v, _ := DecryptString(encryptedString, password)
			return v
		},
	}

	templateText, err := ioutil.ReadFile(templatefile)
	if err != nil {
		return err
	}

	tpl, err := template.New(templatefile).Option("missingkey=error").Funcs(funcMap).Parse(string(templateText))
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	for _, v := range os.Environ() {
		pair := strings.Split(v, "=")
		vars[pair[0]] = pair[1]
	}

	vars, err = DecryptValues(vars, password)
	if err != nil {
		return err
	}

	out := os.Stdout
	if outfile != "" {
		out, err = os.Create(outfile)
		if err != nil {
			return err
		}
	}
	defer out.Close()

	return tpl.Execute(out, vars)
}
