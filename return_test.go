package citest

import (
	"testing"
)

func TestReturnString(t *testing.T) {
	if returnString() != "ABC" {
		t.Errorf("it expected 'ABC' but %v", returnString())
	}
}
