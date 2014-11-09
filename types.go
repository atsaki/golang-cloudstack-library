package cloudstack

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
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
	var i interface{}
	nb.Bool = false
	nb.Valid = false

	if err = json.Unmarshal(b, &i); err != nil {
		return err
	}

	switch v := i.(type) {
	case bool:
		nb.Bool = v
		nb.Valid = true
	case string:
		v = strings.ToLower(v)
		switch v {
		case "true":
			nb.Bool = true
			nb.Valid = true
		case "false":
			nb.Bool = false
			nb.Valid = true
		default:
			return fmt.Errorf("Can't convert to bool: %s", v)
		}
	}
	return nil
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

func (id *ID) MarshalJSON() ([]byte, error) {
	if id.Valid {
		return json.Marshal(id.String)
	} else {
		return json.Marshal(nil)
	}
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
