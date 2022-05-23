package sprintsd_test

import (
	"testing"

	"github.com/royge/sprintsd"
)

func Test_ExtractLocation(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "projects/1234567890/regions/asia-southeast1",
			want:  "asia-southeast1",
		},
		{
			input: "projects/1234567890/regions/asia-southeast2",
			want:  "asia-southeast2",
		},
		{
			input: "projects/1234567890/regions/",
			want:  "",
		},
		{
			input: "projects/1234567890",
			want:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := sprintsd.ExtractLocation(tc.input)

			if tc.want != got {
				t.Errorf("want location `%v`, got `%v`", tc.want, got)
			}
		})
	}
}
