package apollo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	Map interface {
		Put(key string, value interface{})
		Get(key string) (Entry, bool)
		Remove(key string)
	}

	Entry struct {
		Key   string
		Value interface{}
	}

	Storage map[string]Entry
)

var _ Map = (*Storage)(nil)

var nilEntry Entry

func (s *Storage) Put(key string, value interface{}) {
	m := *s

	e, existed := m[key]
	if existed {
		e.Value = value
	} else {
		e = Entry{Key: key, Value: value}
	}

	m[key] = e
}

func (s *Storage) Get(key string) (Entry, bool) {
	m := *s

	e, existed := m[key]
	if existed {
		return e, true
	}

	return nilEntry, false
}

func (s *Storage) Remove(key string) {
	m := *s
	delete(m, key)
}

func (e Entry) GetByKindOrNil(k reflect.Kind) interface{} {
	switch k {
	case reflect.String:

	}
}

func (e Entry) StringDefault(def string) string {
	v := e.Value
	if v == nil {
		return def
	}

	if vString, ok := v.(string); ok {
		return vString
	}

	val := fmt.Sprintf("%v", v)
	if val != "" {
		return val
	}

	return def
}

func (e Entry) String() string {
	return e.StringDefault("")
}

func (e Entry) StringTrim() string {
	return strings.TrimSpace(e.String())
}

var parseErrMsg = "unable to parse the %s with key: %s"

func (e Entry) IntDefault(def int) (int, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "int", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.Atoi(vt)
		if err != nil {
			return def, err
		}
		return val, nil
	case int:
		return vt, nil
	case int8:
		return int(vt), nil
	case int16:
		return int(vt), nil
	case int32:
		return int(vt), nil
	case int64:
		return int(vt), nil
	case uint:
		return int(vt), nil
	case uint8:
		return int(vt), nil
	case uint16:
		return int(vt), nil
	case uint32:
		return int(vt), nil
	case uint64:
		return int(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "int", e.Key)
}

func (e Entry) Int8Default(def int8) (int8, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "int8", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseInt(vt, 10, 8)
		if err != nil {
			return def, err
		}
		return int8(val), nil
	case int:
		return int8(vt), nil
	case int8:
		return vt, nil
	case int16:
		return int8(vt), nil
	case int32:
		return int8(vt), nil
	case int64:
		return int8(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "int8", e.Key)
}

func (e Entry) Int16Default(def int16) (int16, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "int16", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseInt(vt, 10, 16)
		if err != nil {
			return def, err
		}
		return int16(val), nil
	case int:
		return int16(vt), nil
	case int8:
		return int16(vt), nil
	case int16:
		return vt, nil
	case int32:
		return int16(vt), nil
	case int64:
		return int16(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "int16", e.Key)
}

func (e Entry) Int32Default(def int32) (int32, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "int32", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseInt(vt, 10, 32)
		if err != nil {
			return def, err
		}
		return int32(val), nil
	case int:
		return int32(vt), nil
	case int8:
		return int32(vt), nil
	case int16:
		return int32(vt), nil
	case int32:
		return vt, nil
	case int64:
		return int32(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "int32", e.Key)
}

func (e Entry) Int64Default(def int64) (int64, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "int64", e.Key)
	}

	switch vt := v.(type) {
	case string:
		return strconv.ParseInt(vt, 10, 64)
	case int64:
		return vt, nil
	case int32:
		return int64(vt), nil
	case int8:
		return int64(vt), nil
	case int:
		return int64(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "int64", e.Key)
}
