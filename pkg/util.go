package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// FileExists returns true if a file exists
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// LoadYAMLFile parses a YAML file to a string map
func LoadYAMLFile(filename string) (map[string]interface{}, error) {
	var values map[string]interface{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return values, err
	}

	err = yaml.Unmarshal(data, &values)
	if err != nil {
		return values, fmt.Errorf("Failed to parse standard input: %v", err)
	}
	return values, nil
}

// WriteYAMLFile writes a string map to a YAML file
func WriteYAMLFile(values map[string]interface{}, filename string) error {
	data, err := yaml.Marshal(&values)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadEnvFile parses a key=value format file to a string map
func LoadEnvFile(filename string) (map[string]interface{}, error) {
	values := make(map[string]interface{})

	file, err := os.Open(filename)
	if err != nil {
		return values, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		kvpair := strings.SplitN(scanner.Text(), "=", 2)
		if len(kvpair) == 2 {
			values[kvpair[0]] = kvpair[1]
		}
	}

	return values, scanner.Err()
}

// WriteEnvFile writes a string map to a key=value format file
func WriteEnvFile(values map[string]interface{}, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for k, v := range values {
		if _, err := f.WriteString(fmt.Sprintf("%s=%s\n", k, v)); err != nil {
			return err
		}
	}
	return nil
}
