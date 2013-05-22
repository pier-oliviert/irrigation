package helpers

import (
	"testing"
)

func TestEmptyStringConversionToInt32(t *testing.T) {
	value := Int32ValueFrom("", -1)
	if value != -1 {
		t.Fail()
		t.Logf("Should have returned -1. Returned %v", value)
	}
}
