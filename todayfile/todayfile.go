package todayfile

import (
	"os"
	"time"
)

var now = time.Now()

// Create a file with the name of today's date.
func Create() {
	os.Create(today() + ".txt")
}

// return the date of today.
func today() string {
	t := now
	const layout = "2006-01-02"

	return t.Format(layout)
}
