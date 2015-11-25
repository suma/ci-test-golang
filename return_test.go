package ci_test

import (
	"testing"
)

func returnString() string {
	return "abc"
}

func TestReturnString(t *testing.T) {
	if returnString() != "ABC" {
		t.Errorf("it expected 'ABC' but %v", returnString())
	}
}
