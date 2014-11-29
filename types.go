package cloudstack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"code.google.com/p/go-uuid/uuid"
)

type APIParameter interface{}

type Resource interface {
	// refresh information about the resource
	Refresh() (Resource, error)
	// update the resource with specified args
	Update(args map[string]interface{}) (Resource, error)
	// delete the resource
	Delete() error
}

type Nullable interface {
	Value() interface{}
	String() string
	IsNil() bool
}

type Setter interface {
	Set(interface{}) error
}

func unmarshalJSON(s Setter, b []byte) error {
	var v interface{}

	// initialize by nil
	if err := s.Set(nil); err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	if err := decoder.Decode(&v); err != nil {
		return err
	}

	return s.Set(v)
}

// Base struct of Nullable
type NullBase struct {
	valid bool
	value interface{}
}

func (n NullBase) MarshalJSON() ([]byte, error) {
	if n.IsNil() {
		return json.Marshal(nil)
	}
	return json.Marshal(n.value)
}

func (nb *NullBase) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(nb, b)
}

// Set value. If nil is given, value is cleared.
func (nb *NullBase) Set(value interface{}) error {

	nb.valid = false
	nb.value = nil

	if value == nil {
		return nil
	}

	nb.valid = true
	nb.value = value

	return nil
}

// Return Value. If no value is set, return nil.
func (nb NullBase) Value() interface{} {
	if nb.IsNil() {
		return nil
	}
	return nb.value
}

// Return Value as String. If no value is set, return "null".
func (nb NullBase) String() string {
	if nb.IsNil() {
		return "null"
	}
	return fmt.Sprint(nb.value)
}

// Check if value is nil.
func (nb NullBase) IsNil() bool {
	return !nb.valid
}

// Nullable Bool
type NullBool struct {
	NullBase
}

func (nb *NullBool) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(nb, b)
}

// Set Value. Value is converted by strconv.ParseBool
func (nb *NullBool) Set(value interface{}) error {

	nb.valid = false
	nb.value = false

	if value == nil {
		return nil
	}

	b, err := strconv.ParseBool(fmt.Sprint(value))
	if err != nil {
		return err
	}

	nb.valid = true
	nb.value = b

	return nil
}

// Return Value as bool
func (nb NullBool) Bool() bool {
	return nb.value.(bool)
}

// Nullable String
type NullString struct {
	NullBase
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(ns, b)
}

// Set Value. Value is converted by fmt.Sprint
func (ns *NullString) Set(value interface{}) error {

	ns.valid = false
	ns.value = ""

	if value == nil {
		return nil
	}

	ns.valid = true
	ns.value = fmt.Sprint(value)

	return nil
}

// Nullable Number
// Value is stored as string.
type NullNumber struct {
	NullString
}

func (nn *NullNumber) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(nn, b)
}

func (nn NullNumber) MarshalJSON() ([]byte, error) {
	return []byte(nn.String()), nil
}

// Set Value. Value is converted to string by fmt.Sprint
func (nn *NullNumber) Set(value interface{}) error {

	nn.valid = false
	nn.value = ""

	if value == nil {
		return nil
	}

	s := fmt.Sprint(value)
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	nn.valid = true
	nn.value = s

	return nil
}

// Return Value as int64
func (nn NullNumber) Int64() (int64, error) {
	if nn.IsNil() {
		return 0, errors.New("NullNumber is nil")
	}
	return strconv.ParseInt(nn.value.(string), 10, 64)
}

// Return Value as uint64
func (nn NullNumber) UInt64() (uint64, error) {
	if nn.IsNil() {
		return 0, errors.New("NullNumber is nil")
	}
	return strconv.ParseUint(nn.value.(string), 10, 64)
}

// Return Value as float64
func (nn NullNumber) Float64() (float64, error) {
	if nn.IsNil() {
		return 0, errors.New("NullNumber is nil")
	}
	return strconv.ParseFloat(nn.value.(string), 64)
}

// UUID or Integer ID
type ID struct {
	NullString
}

func (id *ID) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(id, b)
}

// Set Value
func (id *ID) Set(value interface{}) error {

	id.valid = false
	id.value = ""

	if value == nil {
		return nil
	}

	s := fmt.Sprint(value)
	uuid := uuid.Parse(s)
	_, err := strconv.ParseFloat(s, 64)

	if uuid != nil || err == nil {
		id.valid = true
		id.value = s
	}

	return nil
}

func (id ID) UUID() uuid.UUID {
	return uuid.Parse(id.String())
}
