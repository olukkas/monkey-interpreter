package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	asserts := assert.New(t)

	hello1 := NewString("Hello World")
	hello2 := NewString("Hello World")

	diff1 := NewString("My name is Jhon")
	diff2 := NewString("My name is Jhon")

	asserts.Equal(
		hello1.HashKey(),
		hello2.HashKey(),
		"strings with same content have different hash keys",
	)

	asserts.Equal(
		diff1.HashKey(),
		diff2.HashKey(),
		"strings with same content have different hash keys",
	)

	asserts.NotEqual(
		hello1.HashKey(),
		diff1.HashKey(),
		"strings with different content have same hash keys",
	)
}

func TestIntegerHashKey(t *testing.T) {
	one := NewInteger(1)
	two := NewInteger(2)

	assert.NotEqual(
		t,
		one.HashKey(),
		two.HashKey(),
		"integers with different content have same hash keys",
	)
}

func TestBooleanHashKey(t *testing.T) {
	bTrue := NewBoolean(true)
	bFalse := NewBoolean(false)

	assert.NotEqual(
		t,
		bTrue.HashKey(),
		bFalse.HashKey(),
		"booleans with different content have same hash keys",
	)
}
