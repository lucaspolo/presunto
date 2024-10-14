package service

import (
	"context"
	"testing"
)

func TestGetQCode(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Call GetQCode with a known value
	meaning, err := GetQCode(ctx, "QAP")

	if err != nil {
		t.Errorf("GetQCode failed: %v", err)
	}

	if meaning != "Está na escuta?" {
		t.Errorf("Expected 'Está na escuta?' but got '%s'", meaning)
	}
}
