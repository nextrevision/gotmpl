package pkg

import "testing"

func TestFileExists(t *testing.T) {
	if !FileExists("../examples/template.tmpl") {
		t.Fatal("File should exist")
	}

	if FileExists("./template.tmpl") {
		t.Fatal("File should not exist")
	}
}

func TestLoadYAMLFile(t *testing.T) {
	vars, err := LoadYAMLFile("../examples/vars.yml")
	if err != nil {
		t.Fatal(err)
	}

	if vars["key1"] != "value1" {
		t.Fatal("Vars have incorrect values")
	}
}

func TestWriteYAMLFile(t *testing.T) {
	vars, err := LoadYAMLFile("../examples/vars.yml")
	if err != nil {
		t.Fatal(err)
	}

	err = WriteYAMLFile(vars, "../examples/vars.yml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadEnvFile(t *testing.T) {
	vars, err := LoadEnvFile("../examples/vars.env")
	if err != nil {
		t.Fatal(err)
	}

	if vars["key1"] != "value1" {
		t.Fatal("Vars have incorrect values")
	}
}

func TestWriteEnvFile(t *testing.T) {
	vars, err := LoadEnvFile("../examples/vars.env")
	if err != nil {
		t.Fatal(err)
	}

	err = WriteEnvFile(vars, "../examples/vars.env")
	if err != nil {
		t.Fatal(err)
	}
}
