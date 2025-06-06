package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size=1",
			size: 1,
		},
		{
			name: "size=5",
			size: 5,
		},
		{
			name: "size=25",
			size: 25,
		},
		{
			name: "size=100",
			size: 100,
		},
		{
			name: "size=1000",
			size: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			str1 := NewRandomString(tt.size)
			str2 := NewRandomString(tt.size)

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size)

			assert.NotEqual(t, str1, str2)
		})
	}
}
