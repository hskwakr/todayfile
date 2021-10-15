package todayfile

import (
	"fmt"
	"os"
	"time"
)

var now = time.Now()

// Create a file with the name of today's date.
func Create() error {
	if _, err := os.Create(today() + ".txt"); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}

// return the date of today.
func today() string {
	const layout = "2006-01-02"

	t := now

	return t.Format(layout)
}
