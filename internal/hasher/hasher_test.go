package hasher

import (
	"testing"
)

var password = "somepass"
var difpassword = "somediffpass"
var hash, err = HashPassword(password)

func TestGenerateFromPassword(t *testing.T) {
	if err != nil {
		t.Fatalf("Bcrypt GenerateFromPassword error: %s", err)
	}
}

func TestPassMatching(t *testing.T) {
	match := CheckPasswordHash(password, hash)
	if !match {
		t.Fatalf("Password doesn't match with hash: %v : %v", password, hash)
	}
}
func TestPassNotMatching(t *testing.T) {
	match := CheckPasswordHash(difpassword, hash)
	if match {
		t.Fatalf("Password shouldn't match with hash: %v : %v", password, hash)
	}
}
