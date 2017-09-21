package pkg

import "testing"

const templateFile = "../examples/template.tmpl"
const outFile = "../examples/template.out"

func TestConvertTemplate(t *testing.T) {
	vars := make(map[string]interface{})
	vars["key1"] = "value1"
	vars["encKey1"] = "ENC|ZGIzMjFmMjAzYTE0MDg5MjJkOGZhMTQ5Ojk1ZmI0ZWZmMWZjZWRhM2MyZGZhNjQyODExNmJmZDJlMzhjNzUyMTYzNzNiM2IyNDdi"

	err := ConvertTemplate(templateFile, vars, outFile, password)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvertTemplateNoFile(t *testing.T) {
	vars := make(map[string]interface{})
	vars["key1"] = "value1"
	vars["encKey1"] = "ENC|ZGIzMjFmMjAzYTE0MDg5MjJkOGZhMTQ5Ojk1ZmI0ZWZmMWZjZWRhM2MyZGZhNjQyODExNmJmZDJlMzhjNzUyMTYzNzNiM2IyNDdi"

	err := ConvertTemplate("./template.tmpl", vars, outFile, password)
	if err == nil {
		t.Fatal("ConvertTemplate should fail when missing template file")
	}
}

func TestConvertTemplateMissingVars(t *testing.T) {
	vars := make(map[string]interface{})
	vars["key1"] = "value1"

	err := ConvertTemplate(templateFile, vars, outFile, password)
	if err == nil {
		t.Fatal("ConvertTemplate should fail when missing template file")
	}
}
