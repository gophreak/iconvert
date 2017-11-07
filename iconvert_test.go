package iconvert_test

import (
	"testing"

	"github.com/gophreak/iconvert"
	"github.com/stretchr/testify/assert"
)

func TestToStringWithString(t *testing.T) {
	value := "1"

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "1", v)
}

func TestToStringWithInt(t *testing.T) {
	value := int(6)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "6", v)
}

func TestToStringWithInt8(t *testing.T) {
	value := int8(127)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "127", v)
}

func TestToStringWithInt16(t *testing.T) {
	value := int16(32767)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "32767", v)
}

func TestToStringWithInt32(t *testing.T) {
	value := int32(2147483647)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "2147483647", v)
}

func TestToStringWithInt64(t *testing.T) {
	value := int64(9223372036854775807)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "9223372036854775807", v)
}

func TestToStringWithFloat32(t *testing.T) {
	value := float32(359.899993896484375)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "359.9", v)
}

func TestToStringWithFloat64(t *testing.T) {
	value := float64(359.8999999999999772626324556767940521240234375)

	v, e := iconvert.ToString(value)
	assert.NoError(t, e)
	assert.Equal(t, "359.9", v)
}

func TestToStringWithStruct(t *testing.T) {
	value := struct{ A int }{A: 5}

	v, e := iconvert.ToString(value)
	assert.Error(t, e)
	assert.Equal(t, e.Error(), "Value struct { A int } is type struct")
	assert.Equal(t, "", v)
}

func TestToFloatWithString(t *testing.T) {
	value := "1"

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(1), v)
}

func TestToFloatWithInt(t *testing.T) {
	value := int(6)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(6), v)
}

func TestToFloatWithInt8(t *testing.T) {
	value := int8(127)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(127), v)
}

func TestToFloatWithInt16(t *testing.T) {
	value := int16(32767)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(32767), v)
}

func TestToFloatWithInt32(t *testing.T) {
	value := int32(2147483647)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(2147483647), v)
}

func TestToFloatWithInt64(t *testing.T) {
	value := int64(9223372036854775807)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(9223372036854775807), v)
}

func TestToFloatWithFloat32(t *testing.T) {
	value := float32(359.9)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(359.9), v)
}

func TestToFloatWithFloat64(t *testing.T) {
	value := float64(179769313.4862932)

	v, e := iconvert.ToFloat(value)
	assert.NoError(t, e)
	assert.Equal(t, float64(179769313.4862932), v)
}

func TestToFloatWithStruct(t *testing.T) {
	value := struct{ A int }{A: 5}

	v, e := iconvert.ToFloat(value)
	assert.Error(t, e)
	assert.Equal(t, e.Error(), "Value struct { A int } is type struct")
	assert.Equal(t, float64(0), v)
}

func TestToFloatWithWord(t *testing.T) {
	value := "Hello"

	v, e := iconvert.ToFloat(value)
	assert.Error(t, e)
	assert.Contains(t, e.Error(), "invalid syntax")
	assert.Equal(t, float64(0), v)
}

func TestToIntWithString(t *testing.T) {
	value := "1"

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(1), v)
}

func TestToIntWithInt(t *testing.T) {
	value := int(6)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(6), v)
}

func TestToIntWithInt8(t *testing.T) {
	value := int8(127)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(127), v)
}

func TestToIntWithInt16(t *testing.T) {
	value := int16(32767)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(32767), v)
}

func TestToIntWithInt32(t *testing.T) {
	value := int32(2147483647)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(2147483647), v)
}

func TestToIntWithInt64(t *testing.T) {
	value := int64(9223372036854775807)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(9223372036854775807), v)
}

func TestToIntWithFloat32(t *testing.T) {
	value := float32(359.9)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(359), v)
}

func TestToIntWithFloat64(t *testing.T) {
	value := float64(179769313.4862932)

	v, e := iconvert.ToInt(value)
	assert.NoError(t, e)
	assert.Equal(t, int64(179769313), v)
}

func TestToIntWithStruct(t *testing.T) {
	value := struct{ A int }{A: 5}

	v, e := iconvert.ToInt(value)
	assert.Error(t, e)
	assert.Equal(t, e.Error(), "Value struct { A int } is type struct")
	assert.Equal(t, int64(0), v)
}

func TestToIntWithWord(t *testing.T) {
	value := "Hello"

	v, e := iconvert.ToInt(value)
	assert.Error(t, e)
	assert.Contains(t, e.Error(), "invalid syntax")
	assert.Equal(t, int64(0), v)
}
