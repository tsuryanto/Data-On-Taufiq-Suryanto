package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("Test Add() with single value", func(t *testing.T) {
		result := Add(3)
		assert.Equal(t, float64(3), result)
	})

	t.Run("Test Add() with decimal and negative value", func(t *testing.T) {
		result := Add(-10, 2.5, 3, 4)
		assert.Equal(t, -0.5, result)
	})

	t.Run("Test Add() with 0 value", func(t *testing.T) {
		result := Add(0, 0, 1, 0)
		assert.Equal(t, float64(1), result)
	})
}

func TestSubstract(t *testing.T) {
	t.Run("Test Substract() with decimal and negative value", func(t *testing.T) {
		result := Substract(-10, 2, -12, 4.5)
		assert.Equal(t, -4.5, result)
	})

	t.Run("Test Substract() with 0 value", func(t *testing.T) {
		result := Substract(0, 0, 1, 0)
		assert.Equal(t, float64(-1), result)
	})
}

func TestDiv(t *testing.T) {
	t.Run("Test Div() with normal value", func(t *testing.T) {
		result := Div(10, 2)
		assert.Equal(t, float64(5), result)
	})

	t.Run("Test Div() with negative value", func(t *testing.T) {
		result := Div(-10, 2)
		assert.Equal(t, float64(-5), result)
	})

	t.Run("Test Div() with 0 value", func(t *testing.T) {
		result := Div(0, 20)
		assert.Equal(t, float64(0), result)
	})

	t.Run("Test Div() by zero", func(t *testing.T) {
		result := Div(20, 0)
		assert.NotEqual(t, float64(0), result)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Test Multiply() with negative value", func(t *testing.T) {
		result := Multiply(-3, 3, -3)
		assert.Equal(t, float64(27), result)
	})

	t.Run("Test Multiply() with decimal value", func(t *testing.T) {
		result := Multiply(2.5, 2, 0.5)
		assert.Equal(t, 2.5, result)
	})

	t.Run("Test Multiply() with 0 value", func(t *testing.T) {
		result := Multiply(2000, 192, 0, 3)
		assert.Equal(t, float64(0), result)
	})
}
