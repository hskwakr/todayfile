package todayfile

import (
	"fmt"
	"os"
	"time"
)

// Create a file with the name of today's date.
func Create() error {
	name := today(time.Now()) + ".txt"

	// if there is a file, do nothing.
	if _, err := os.Stat(name); err == nil {
		return nil
	}

	if _, err := os.Create(name); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}

// return the date of today.
func today(t time.Time) string {
	const layout = "2006-01-02"

	return t.Format(layout)
}
