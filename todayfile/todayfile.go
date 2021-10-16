package todayfile

import (
	"fmt"
	"os"
	"time"
)

// Create a file with the name of today's date.
func Create(name string) error {
	// if the file already exists, do nothing.
	if _, err := os.Stat(name); err == nil {
		return nil
	}

	if _, err := os.Create(name); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}

// Date of the given time.
func Date(t time.Time) string {
	const layout = "2006-01-02"

	return t.Format(layout)
}
