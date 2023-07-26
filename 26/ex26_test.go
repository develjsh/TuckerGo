package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

 func TestSquare1(t *testing.T) {
	assert.Equal(t, 81, square(9), "square(9) should be 81")
	// rst := square(9)
	// if rst != 81 {
	// 	t.Errorf("square(9) should be 81 but returns %d", rst)
	// }
}

func TestSquare2(t *testing.T) {
	rst := square(9)
	if rst != 8 {
		t.Errorf("square(9) should be 81 but returns %d", rst)
	}
}