package password

import (
	"testing"
)

func TestPassowrd(t *testing.T) {
	pwd := "my password"
	code := Encrypt(pwd)
	pass := Verify("error password", code)
	if pass {
		t.Error("error verify")
	}
	pass = Verify(pwd, code)
	if !pass {
		t.Error("error verify")
	}
}
