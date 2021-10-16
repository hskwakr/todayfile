package todayfile

import (
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
