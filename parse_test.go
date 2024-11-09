package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestParseTitle(t *testing.T) {
	r, _ := os.Open("fixtures/program_day.html")
	type args struct {
		r            io.Reader
		programTime  string
		programTitle string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{
				r:            nil,
				programTime:  "",
				programTitle: "",
			},
			want: "",
		},
		{
			name: "empty reader",
			args: args{
				r:            bytes.NewReader([]byte{}),
				programTime:  DEFAULT_TIME,
				programTitle: DEFAULT_TITLE,
			},
			want: "",
		},
		{
			name: "happy path",
			args: args{
				r:            r,
				programTime:  DEFAULT_TIME,
				programTitle: DEFAULT_TITLE,
			},
			want: "Das ist ein \"Test\"",
		},
		{
			name: "time not found",
			args: args{
				r:            r,
				programTime:  "12:34",
				programTitle: DEFAULT_TITLE,
			},
			want: "",
		},
		{
			name: "title not found",
			args: args{
				r:            r,
				programTime:  DEFAULT_TIME,
				programTitle: "Das ist ein anderer Test",
			},
			want: "",
		},
		{
			name: "nothing found",
			args: args{
				r:            r,
				programTime:  "12:34",
				programTitle: "Das ist ein anderer Test",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTitle(tt.args.r, tt.args.programTime, tt.args.programTitle); got != tt.want {
				t.Errorf("ParseTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
