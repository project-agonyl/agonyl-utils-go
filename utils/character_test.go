package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClassName(t *testing.T) {
	tests := []struct {
		name  string
		class byte
		want  string
	}{
		{"Holy Knight", 1, "Holy Knight"},
		{"Mage", 2, "Mage"},
		{"Archer", 3, "Archer"},
		{"Warrior default for zero", 0, "Warrior"},
		{"Warrior for unknown", 4, "Warrior"},
		{"Warrior for unknown", 255, "Warrior"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetClassName(tt.class)
			assert.Equal(t, tt.want, got)
		})
	}
}
