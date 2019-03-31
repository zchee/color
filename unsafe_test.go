package color

import (
	"reflect"
	"testing"
)

func Test_unsafeToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []byte
	}{
		{
			name: "hello",
			s:    "Hello world",
			want: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
		},
		{
			name: "empty",
			s:    "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unsafeToSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unsafeToSlice(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
