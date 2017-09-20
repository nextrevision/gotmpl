package pkg

import "testing"

func TestNewPassword(t *testing.T) {
	password := NewPassword(32)
	if len(password) != 32 {
		t.Fatal("Password length does not match")
	}
}
