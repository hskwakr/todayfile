package todayfile

import (
	"os"
	"time"
)

var now = time.Now()

// Create a file with the name of today's date.
func Create() error {
	_, err := os.Create(today() + ".txt")

	return err
}

// return the date of today.
func today() string {
	const layout = "2006-01-02"

	t := now

	return t.Format(layout)
}
