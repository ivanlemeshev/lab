package dtid_test

import (
	"dtid"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	// Simple test to check if the function returns the current time in the
	// format "20060102150405". However, this test might fail if the test runs
	// at the turn of the second
	expected := time.Now().UTC().Format("20060102150405")
	result := dtid.New()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
