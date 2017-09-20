package pkg

import "testing"

const password = "abcdefghijklmnopqrstuvwxyz012345"

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
	vars["enckey1"] = "ENC|YjE0YzY4ZmExOGQ4NWI5MDljNGI5M2Q5OmM3MDk5ZDQzNjM5NTY3ZTU0NzMwZjg0YjZiNzhiOTg3NDBmZTFkYWY2ZGZlODUyZWZm"
	vars["key2"] = "value2"
	vars["enckey2"] = "ENC|OTYxYjA0OWJmZWI2NDE1OWRiNWZiYjdlOjUxNTI2YzI4ZDI3Njk0Yjk2YTMzMDNlZGJmMGUyZjYzNTU1NzA3NjkxZjJhZTY5MzI4"

	result, err := DecryptValues(vars, password)
	if err != nil {
		t.Fatal(err)
	}

	if result["key1"] != "value1" ||
		result["key2"] != "value2" ||
		result["enckey1"] != "encvalue1" ||
		result["enckey2"] != "encvalue2" {
		t.Fatalf("Unencrypted values do not match: %+v", result)
	}
}
