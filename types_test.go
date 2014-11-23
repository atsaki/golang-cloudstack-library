package cloudstack

import (
	"testing"
)

func testNullBool(t *testing.T, v interface{}, expected bool) error {

	nb := new(NullBool)

	if err := nb.Set(v); err != nil {
		return err
	}
	if nb.IsNil() {
		t.Errorf("input: %v, isNil", v)
		return nil
	}

	if nb.Bool() != expected {
		t.Errorf("input: %v, actual: %v, expected: %v", v, nb.Bool(), expected)
	}
	return nil
}

func TestNullBool(t *testing.T) {
	testNullBool(t, true, true)
	testNullBool(t, false, false)
	testNullBool(t, 1, true)
	testNullBool(t, 0, false)
	testNullBool(t, "true", true)
	testNullBool(t, "false", false)
	testNullBool(t, "", false)
}

func testNullString(t *testing.T, v interface{}, expected string) error {

	ns := new(NullString)

	if err := ns.Set(v); err != nil {
		return err
	}
	if ns.IsNil() {
		t.Errorf("input: %v, isNil", v)
		return nil
	}

	if ns.String() != expected {
		t.Errorf("input: %v, actual: %v, expected: %v", v, ns.String(), expected)
	}
	return nil
}

func TestNullString(t *testing.T) {
	testNullString(t, "", "")
	testNullString(t, "string", "string")
	testNullString(t, 0, "0")
	testNullString(t, true, "true")
}

func testNullNumber(t *testing.T, v interface{}, expected string) error {

	nn := new(NullNumber)

	if err := nn.Set(v); err != nil {
		return err
	}
	if nn.IsNil() {
		t.Errorf("input: %v, isNil", v)
		return nil
	}

	if nn.String() != expected {
		t.Errorf("input: %v, actual: %v, expected: %v", v, nn.String(), expected)
	}
	return nil
}

func TestNullNumber(t *testing.T) {
	testNullNumber(t, "0", "0")
	testNullNumber(t, 0, "0")
	testNullNumber(t, 100, "100")
	testNullNumber(t, 10000000000000, "10000000000000")
	testNullNumber(t, -3.14, "-3.14")
}
