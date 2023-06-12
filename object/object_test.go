package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := NewString("Hello World")
	hello2 := NewString("Hello World")

	diff1 := NewString("My name is Jhon")
	diff2 := NewString("My name is Jhon")

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one := NewInteger(1)
	two := NewInteger(2)

	if one.HashKey() == two.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	bTrue := NewBoolean(true)
	bFalse := NewBoolean(false)

	if bTrue.HashKey() == bFalse.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}
