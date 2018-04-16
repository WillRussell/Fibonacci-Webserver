package main

import "testing"


func TestParseUserNumberInput(t *testing.T) {

	t.Log("Parse out url parameters to isolate integer used in fibonacci formula (expected num: 123)")

	actualResult := parseInput("/123")
	expectedResult := 123

	if actualResult != expectedResult {
		t.Errorf("Expected integer 123, but it was %d instead.", actualResult)
	}
}