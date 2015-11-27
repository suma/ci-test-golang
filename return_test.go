package citest

import (
	"testing"
)

func TestReturnInt(t *testing.T) {
	if returnInt() != 100 {
		t.Errorf("it expected 100 but %v", returnInt())
	}
}

func TestReturnString(t *testing.T) {
	if returnString() != "ABC" {
		t.Errorf("it expected 'ABC' but %v", returnString())
	}
}
