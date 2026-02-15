package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNationName(t *testing.T) {
	tests := []struct {
		name   string
		nation byte
		want   string
	}{
		{"Quanato", 1, "Quanato"},
		{"Temoz default for zero", 0, "Temoz"},
		{"Temoz for unknown", 2, "Temoz"},
		{"Temoz for unknown", 255, "Temoz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNationName(tt.nation)
			assert.Equal(t, tt.want, got)
		})
	}
}
