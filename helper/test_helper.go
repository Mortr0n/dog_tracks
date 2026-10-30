package helper

import (
	"encoding/json"
	"testing"
)

func ObjectCheckFormatter[T any](t *testing.T, expected T, actual T) {
	t.Helper()
	expectedJson, err := json.MarshalIndent(expected, "", "  ")
	if err != nil {
		t.Fatalf("error marshaling expected : %v", err)
	}

	tsJson, err := json.MarshalIndent(actual, "", "  ")
	if err != nil {
		t.Fatalf("error marshaling traing session : %v", err)
	}
	t.Errorf("incorrect session, expected: %s, got %s", expectedJson, tsJson)
}
