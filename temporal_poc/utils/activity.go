package utils

import (
	"context"
	"fmt"
)

// SampleActivity is a simple Temporal activity.
func SampleActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
