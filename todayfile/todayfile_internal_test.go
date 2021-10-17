package todayfile

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	tests := []struct {
		name string
		in   time.Time
		want string
	}{
		{
			name: "case 1: Proper",
			in:   time.Date(1999, time.January, 2, 0, 0, 0, 0, time.UTC),
			want: "1999-01-02",
		},
	}

	t.Parallel()

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Date(test.in)
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func createFakeFile(name string, msg string) bool {
	const mode = 0644

	if err := os.WriteFile(name, []byte(msg), mode); err != nil {
		return false
	}

	return true
}

func readFakeFile(name string) string {
	msg, err := os.ReadFile(name)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	return string(msg)
}

func removeFakeFile(name string) bool {
	if err := os.Remove(name); err != nil {
		return false
	}

	return true
}

func TestCreate_ShouldNotOverwriteFile(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "case 1: Proper",
			in:   "Hello",
			want: "Hello",
		},
	}

	fn := "1992-01-02.txt"

	t.Parallel()

	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if success := createFakeFile(fn, test.in); !success {
				t.Errorf("failed to create fake file")
			}

			if err := Create(fn); err != nil {
				t.Errorf("%v", err)
			}

			got := readFakeFile(fn)

			if success := removeFakeFile(fn); !success {
				t.Errorf("failed to remove fake file")
			}

			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
