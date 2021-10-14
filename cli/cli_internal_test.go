package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "case 1: Proper",
			in:   ApplicationName,
			want: 0,
		},
	}

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	app := &CLI{
		OutStream: outStream,
		ErrStream: errStream,
	}

	t.Parallel()
	for _, v := range tests {
		test := v

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			args := strings.Split(test.in, " ")
			got := app.parse(args)

			if got != test.want {
				t.Errorf("ExitStatus: %d, want: %d",
					got,
					test.want,
				)
			}
		})
	}
}
