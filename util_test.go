package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllTrue(t *testing.T) {
	//t.Skip("skipping test")
	s := []bool{true, true, true, true, true, true}
	t.Run("Should return true", func(t *testing.T) {
		assert.True(t, allTrue(s, len(s)))
	})
	s = []bool{true, true, true, true, false, true}
	t.Run("Should return false", func(t *testing.T) {
		assert.False(t, allTrue(s, len(s)))
	})
}

func TestUpdate(t *testing.T) {
	//t.Skip("skipping test")
	s1 := []float32{3.6, 8.7, 1.8, 18.2}
	s2 := make([]float32, len(s1))
	update(s2, s1, len(s1))
	t.Run("Should return true", func(t *testing.T) {
		assert.Equal(t, s1, s2)
	})
}

func TestConvertLevel(t *testing.T) {
	//t.Skip("skipping test")
	res := convertLevel(0.02)
	t.Run("Should be equal", func(t *testing.T) {
		assert.Equal(t, float32(-34), res)
	})
	res = convertLevel(-0.02)
	t.Run("Should be equal", func(t *testing.T) {
		assert.Equal(t, float32(-200), res)
	})
}
