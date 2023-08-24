package utils

import "testing"

func TestRandomString(t *testing.T) {
	str1 := RandomString()
	str2 := RandomString()

	if len(str1) < 1 {
		t.Fatalf("bad random string: should be not empty")
	}

	if str1 == str2 {
		t.Fatalf("bad random string: str1 should not be equals to str2")
	}
}
