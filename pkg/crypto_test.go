package pkg

import "testing"

const password = "password"

func TestEncryptDecryptString(t *testing.T) {
	encResult, err := EncryptString("test", password)
	if err != nil {
		t.Fatal(err)
	}

	result, err := DecryptString(encResult, password)
	if err != nil {
		t.Fatal(err)
	}

	if result != "test" {
		t.Fatal("Unencrypted strings do not match")
	}
}

func TestDecryptVars(t *testing.T) {
	vars := make(map[string]interface{})
	vars["key1"] = "value1"
	vars["key2"] = "value2"
	vars["encKey1"] = "ENC|ZGIzMjFmMjAzYTE0MDg5MjJkOGZhMTQ5Ojk1ZmI0ZWZmMWZjZWRhM2MyZGZhNjQyODExNmJmZDJlMzhjNzUyMTYzNzNiM2IyNDdi"
	vars["encKey2"] = "ENC|NzJlZmY0MDU2NGMwOTU4OWUxYWM1Y2ZkOjAzNWViZDAzY2JmN2I5NWQ3OTNlY2YyN2E5ZjQzYjc1M2JkYWU3YTU4ODlhNDIxZTNh"

	result, err := DecryptValues(vars, password)
	if err != nil {
		t.Fatal(err)
	}

	if result["key1"] != "value1" ||
		result["key2"] != "value2" ||
		result["encKey1"] != "encValue1" ||
		result["encKey2"] != "encValue2" {
		t.Fatalf("Unencrypted values do not match: %+v", result)
	}
}
