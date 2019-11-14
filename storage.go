package apollo

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type (
	Map interface {
		Put(key string, value interface{})
		Get(key string) (Entry, bool)
		Remove(key string)
		Clear()
		Len() int
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

func (s *Storage) Clear() {
	*s = Storage{}
}

func (s *Storage) Len() int {
	m := *s
	return len(m)
}

func (e Entry) GetByKindOrNil(k reflect.Kind) interface{} {
	switch k {
	case reflect.String:
		v := e.StringDefault("__$nf")
		if v == "__$nf" {
			return nil
		}
		return v
	case reflect.Int:
		v, err := e.IntDefault(-1)
		if err != nil {
			return nil
		}
		return v
	case reflect.Int64:
		v, err := e.Int64Default(-1)
		if err != nil {
			return nil
		}
		return v
	case reflect.Bool:
		v, err := e.BoolDefault(false)
		if err != nil {
			return nil
		}
		return v
	default:
		return e.Value
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

func (e Entry) UintDefault(def uint) (uint, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "uint", e.Key)
	}

	x64 := strconv.IntSize == 64
	var maxValue uint64 = math.MaxUint32
	if x64 {
		maxValue = math.MaxUint64
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseUint(vt, 10, strconv.IntSize)
		if err != nil {
			return def, err
		}
		if val > uint64(maxValue) {
			return def, fmt.Errorf(parseErrMsg, "uint", e.Key)
		}
		return uint(val), nil
	case uint:
		return vt, nil
	case uint8:
		return uint(vt), nil
	case uint16:
		return uint(vt), nil
	case uint32:
		return uint(vt), nil
	case uint64:
		if vt > uint64(maxValue) {
			return def, fmt.Errorf(parseErrMsg, "uint", e.Key)
		}
		return uint(vt), nil
	case int:
		if vt < 0 || vt > int(maxValue) {
			return def, fmt.Errorf(parseErrMsg, "uin", e.Key)
		}
		return uint(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "uin", e.Key)
}

func (e Entry) Uint8Default(def uint8) (uint8, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseUint(vt, 10, 8)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		if val > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(val), nil
	case uint:
		if vt > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(vt), nil
	case uint8:
		return vt, nil
	case uint16:
		if vt > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(vt), nil
	case uint32:
		if vt > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(vt), nil
	case uint64:
		if vt > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(vt), nil
	case int:
		if vt < 0 || vt > math.MaxUint8 {
			return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
		}
		return uint8(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "uint8", e.Key)
}

func (e Entry) Uint16Default(def uint16) (uint16, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseUint(vt, 10, 16)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		if val > math.MaxUint16 {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		return uint16(val), nil
	case uint:
		if vt > math.MaxUint16 {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		return uint16(vt), nil
	case uint8:
		return uint16(vt), nil
	case uint16:
		return vt, nil
	case uint32:
		if vt > math.MaxUint16 {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		return uint16(vt), nil
	case uint64:
		if vt > math.MaxUint16 {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		return uint16(vt), nil
	case int:
		if vt < 0 || vt > math.MaxUint16 {
			return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
		}
		return uint16(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "uint16", e.Key)
}

func (e Entry) Uint32Default(def uint32) (uint32, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseUint(vt, 10, 32)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		if val > math.MaxUint32 {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		return uint32(val), nil
	case uint:
		if vt > math.MaxUint32 {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		return uint32(vt), nil
	case uint8:
		return uint32(vt), nil
	case uint16:
		return uint32(vt), nil
	case uint32:
		return vt, nil
	case uint64:
		if vt > math.MaxUint32 {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		return uint32(vt), nil
	case int32:
		if vt < 0 {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		return uint32(vt), nil
	case int64:
		if vt < 0 || vt > math.MaxUint32 {
			return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
		}
		return uint32(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "uint32", e.Key)
}

func (e Entry) Uint64Default(def uint64) (uint64, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseUint(vt, 10, 64)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
		}
		if val > math.MaxUint64 {
			return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
		}
		return val, nil
	case uint:
		return uint64(vt), nil
	case uint8:
		return uint64(vt), nil
	case uint16:
		return uint64(vt), nil
	case uint32:
		return uint64(vt), nil
	case uint64:
		return vt, nil
	case int:
		if vt < 0 {
			return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
		}
		return uint64(vt), nil
	case int64:
		if vt < 0 {
			return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
		}
		return uint64(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "uint64", e.Key)
}

func (e Entry) Float32Default(def float32) (float32, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "float32", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseFloat(vt, 32)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "float32", e.Key)
		}
		if val > math.MaxFloat32 {
			return def, fmt.Errorf(parseErrMsg, "float32", e.Key)
		}
		return float32(val), nil
	case float32:
		return vt, nil
	case float64:
		if vt > math.MaxFloat32 {
			return def, fmt.Errorf(parseErrMsg, "float32", e.Key)
		}
		return float32(vt), nil
	case int:
		return float32(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "float32", e.Key)
}

func (e Entry) Float64Default(def float64) (float64, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "float64", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseFloat(vt, 64)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "float64", e.Key)
		}
		return val, nil
	case float32:
		return float64(vt), nil
	case float64:
		return vt, nil
	case int:
		return float64(vt), nil
	case int64:
		return float64(vt), nil
	case uint32:
		return float64(vt), nil
	case uint64:
		return float64(vt), nil
	}

	return def, fmt.Errorf(parseErrMsg, "float64", e.Key)
}

func (e Entry) BoolDefault(def bool) (bool, error) {
	v := e.Value
	if v == nil {
		return def, fmt.Errorf(parseErrMsg, "bool", e.Key)
	}

	switch vt := v.(type) {
	case string:
		val, err := strconv.ParseBool(vt)
		if err != nil {
			return def, fmt.Errorf(parseErrMsg, "bool", e.Key)
		}
		return val, nil
	case bool:
		return vt, nil
	case int:
		if vt == 0 {
			return false, nil
		}
		return true, nil
	}

	return def, fmt.Errorf(parseErrMsg, "bool", e.Key)
}

func (e Entry) GetValue() interface{} {
	return e.Value
}

func (e Entry) GetKey() string {
	return e.Key
}

func (s *Storage) GetValueDefault(key string, def interface{}) interface{} {
	v, existed := s.Get(key)
	if !existed || v.Value == nil {
		return def
	}
	vv := v.GetValue()
	if vv == nil {
		return def
	}
	return vv
}

func (s *Storage) GetValue(key string) interface{} {
	return s.GetValueDefault(key, nil)
}

func (s *Storage) GetStringDefault(key string, def string) string {
	e, existed := s.Get(key)
	if !existed {
		return def
	}
	return e.StringDefault(def)
}

func (s *Storage) GetString(key string) string {
	return s.GetStringDefault(key, "")
}

func (s *Storage) GetStringTrim(key string) string {
	return strings.TrimSpace(s.GetString(key))
}

func (s *Storage) GetInt(key string) (int, error) {
	e, exist := s.Get(key)
	if !exist {
		return 0, fmt.Errorf(parseErrMsg, "int", key)
	}
	return e.IntDefault(-1)
}

func (s *Storage) GetIntDefault(key string, def int) int {
	if v, err := s.GetInt(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetInt8(key string) (int8, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "int8", key)
	}
	return e.Int8Default(-1)
}

func (s *Storage) GetInt8Default(key string, def int8) int8 {
	if v, err := s.GetInt8(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetInt16(key string) (int16, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "int16", key)
	}
	return e.Int16Default(-1)
}

func (s *Storage) GetInt16Default(key string, def int16) int16 {
	if v, err := s.GetInt16(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetInt32(key string) (int32, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "int32", key)
	}
	return e.Int32Default(-1)
}

func (s *Storage) GetInt32Default(key string, def int32) int32 {
	if v, err := s.GetInt32(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetInt64(key string) (int64, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "int64", key)
	}
	return e.Int64Default(-1)
}

func (s *Storage) GetInt64Default(key string, def int64) int64 {
	if v, err := s.GetInt64(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetUint(key string) (uint, error) {
	e, exist := s.Get(key)
	if !exist {
		return 0, fmt.Errorf(parseErrMsg, "uint", key)
	}
	return e.UintDefault(0)
}

func (s *Storage) GetUintDefault(key string, def uint) uint {
	if v, err := s.GetUint(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetUint8(key string) (uint8, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "uint8", key)
	}
	return e.Uint8Default(0)
}

func (s *Storage) GetUint8Default(key string, def uint8) uint8 {
	if v, err := s.GetUint8(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetUint16(key string) (uint16, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "uint16", key)
	}
	return e.Uint16Default(0)
}

func (s *Storage) GetUint16Default(key string, def uint16) uint16 {
	if v, err := s.GetUint16(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetUint32(key string) (uint32, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "uint32", key)
	}
	return e.Uint32Default(0)
}

func (s *Storage) GetUint32Default(key string, def uint32) uint32 {
	if v, err := s.GetUint32(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetUint64(key string) (uint64, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "uint64", key)
	}
	return e.Uint64Default(0)
}

func (s *Storage) GetUint64Default(key string, def uint64) uint64 {
	if v, err := s.GetUint64(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetFloat32(key string) (float32, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "float32", key)
	}
	return e.Float32Default(-1)
}

func (s *Storage) GetFloat32Default(key string, def float32) float32 {
	if v, err := s.GetFloat32(key); err == nil {
		return v
	}
	return def
}

func (s *Storage) GetFloat64(key string) (float64, error) {
	e, existed := s.Get(key)
	if !existed {
		return 0, fmt.Errorf(parseErrMsg, "float64", key)
	}
	return e.Float64Default(-1)
}

func (s *Storage) GetFloat64Default(key string, def float64) float64 {
	if v, err := s.GetFloat64(key); err == nil {
		return v
	}
	return def
}
