package cloudstack

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullBool struct {
	sql.NullBool
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (nb *NullBool) UnmarshalJSON(b []byte) (err error) {
	val := false
	if err = json.Unmarshal(b, &val); err == nil {
		nb.Bool = val
		nb.Valid = true
	} else {
		nb.Bool = false
		nb.Valid = false
	}
	return err
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (ni *NullInt64) UnmarshalJSON(b []byte) (err error) {
	val := int64(0)
	if err = json.Unmarshal(b, &val); err == nil {
		ni.Int64 = val
		ni.Valid = true
	} else {
		ni.Int64 = 0
		ni.Valid = false
	}
	return err
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (nf *NullFloat64) UnmarshalJSON(b []byte) (err error) {
	val := float64(0)
	if err = json.Unmarshal(b, &val); err == nil {
		nf.Float64 = val
		nf.Valid = true
	} else {
		nf.Float64 = 0
		nf.Valid = false
	}
	return err
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	} else {
		return json.Marshal(nil)
	}
}

func (ns *NullString) UnmarshalJSON(b []byte) (err error) {
	var s interface{}
	if err = json.Unmarshal(b, &s); err == nil {
		ns.String = fmt.Sprint(s)
		ns.Valid = true
	} else {
		ns.String = ""
		ns.Valid = false
	}
	return err
}

type ID struct {
	NullString
}

func (id *ID) UnmarshalJSON(b []byte) (err error) {
	var i interface{}
	if err = json.Unmarshal(b, &i); err == nil {
		switch v := i.(type) {
		case float64:
			id.String = fmt.Sprint(uint64(v))
		default:
			id.String = fmt.Sprint(v)
		}
		id.Valid = true
	} else {
		id.String = ""
		id.Valid = false
	}
	return err
}

type Object interface {
	GetValue()
}
type Boolean struct {
	value bool
}

func (boolean *Boolean) GetValue() bool {
	return boolean.value
}

func (boolean *Boolean) SetValue(b bool) {
	boolean.value = b
}

func (boolean *Boolean) MarshalJSON() (b []byte, err error) {
	return json.Marshal(boolean.GetValue())
}

func (boolean *Boolean) UnmarshalJSON(b []byte) (err error) {
	val := false
	if err = json.Unmarshal(b, &val); err == nil {
		boolean.SetValue(val)
		return nil
	}
	return err
}

type Number struct {
	value float64
}

func (num *Number) GetValue() float64 {
	return num.value
}

func (num *Number) SetValue(n float64) {
	num.value = n
}

func (num *Number) MarshalJSON() (b []byte, err error) {
	return json.Marshal(num.GetValue())
}

func (num *Number) UnmarshalJSON(b []byte) (err error) {
	n := float64(0)
	if err = json.Unmarshal(b, &n); err == nil {
		num.SetValue(n)
		return nil
	}
	return err
}

type Integer struct {
	value int
}

func (num *Integer) GetValue() int {
	return num.value
}

func (num *Integer) SetValue(n int) {
	num.value = n
}

func (num *Integer) MarshalJSON() (b []byte, err error) {
	return json.Marshal(num.GetValue())
}

func (num *Integer) UnmarshalJSON(b []byte) (err error) {
	n := int(0)
	if err = json.Unmarshal(b, &n); err == nil {
		num.SetValue(n)
		return nil
	}
	return err
}

type String struct {
	value string
}

func (str *String) GetValue() string {
	return str.value
}

func (str *String) SetValue(s string) {
	str.value = s
}

func (str *String) MarshalJSON() (b []byte, err error) {
	return json.Marshal(str.GetValue())
}

func (str *String) UnmarshalJSON(b []byte) (err error) {
	var s interface{}
	if err = json.Unmarshal(b, &s); err == nil {
		str.SetValue(fmt.Sprint(s))
		return nil
	}
	return err
}
