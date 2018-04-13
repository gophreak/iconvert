package iconvert_test

import (
	"fmt"
	"testing"

	"github.com/gophreak/iconvert"
	"github.com/stretchr/testify/assert"
)

func messageForFailedTest(index int, message string) string {
	return fmt.Sprintf("test case: [%v] with description: %s", index, message)
}

func TestToString(t *testing.T) {
	cases := []struct {
		input                interface{}
		result               string
		wantError            bool
		errorMessageContains string
		description          string
	}{
		{
			"Hello",
			"Hello",
			false,
			"",
			"convert string that is not a number",
		},
		{
			struct{ A int }{A: 5},
			"",
			true,
			"Value struct { A int } is type struct",
			"convert struct to int",
		},
		{
			float64(359.8999999999999772626324556767940521240234375),
			"359.9",
			false,
			"",
			"convert float64",
		},
		{
			float32(359.899993896484375),
			"359.9",
			false,
			"",
			"convert float32",
		},
		{
			int64(9223372036854775807),
			"9223372036854775807",
			false,
			"",
			"convert int64",
		},
		{
			int32(2147483647),
			"2147483647",
			false,
			"",
			"convert int32",
		},
		{
			int16(32767),
			"32767",
			false,
			"",
			"convert int16",
		},
		{
			int8(127),
			"127",
			false,
			"",
			"convert int8",
		},
		{
			int(6),
			"6",
			false,
			"",
			"convert int",
		},
		{
			"1",
			"1",
			false,
			"",
			"convert string",
		},
	}
	for i, tt := range cases {
		v, e := iconvert.ToString(tt.input)
		if tt.wantError {
			assert.Errorf(t, e, messageForFailedTest(i, tt.description))
			if tt.errorMessageContains != "" && e != nil {
				assert.Containsf(t, e.Error(), tt.errorMessageContains, messageForFailedTest(i, tt.description))
			}

			continue
		}
		assert.NoErrorf(t, e, messageForFailedTest(i, tt.description))
		assert.Equalf(t, tt.result, v, messageForFailedTest(i, tt.description))
	}
}

func TestToFloat(t *testing.T) {
	cases := []struct {
		input                interface{}
		result               float64
		wantError            bool
		errorMessageContains string
		description          string
	}{
		{
			"Hello",
			0.0,
			true,
			"invalid syntax",
			"convert string that is not a number",
		},
		{
			struct{ A int }{A: 5},
			0.0,
			true,
			"Value struct { A int } is type struct",
			"convert struct to int",
		},
		{
			float64(179769313.4862932),
			179769313.4862932,
			false,
			"",
			"convert float64",
		},
		{
			float32(359.9),
			359.9,
			false,
			"",
			"convert float32",
		},
		{
			int64(9223372036854775807),
			9223372036854775807.0,
			false,
			"",
			"convert int64",
		},
		{
			int32(2147483647),
			2147483647.0,
			false,
			"",
			"convert int32",
		},
		{
			int16(32767),
			32767.0,
			false,
			"",
			"convert int16",
		},
		{
			int8(127),
			127.0,
			false,
			"",
			"convert int8",
		},
		{
			int(6),
			6.0,
			false,
			"",
			"convert int",
		},
		{
			"1",
			1.0,
			false,
			"",
			"convert string",
		},
	}
	for i, tt := range cases {
		v, e := iconvert.ToFloat(tt.input)
		if tt.wantError {
			assert.Errorf(t, e, messageForFailedTest(i, tt.description))
			if tt.errorMessageContains != "" {
				assert.Containsf(t, e.Error(), tt.errorMessageContains, messageForFailedTest(i, tt.description))
			}

			continue
		}
		assert.NoErrorf(t, e, messageForFailedTest(i, tt.description))
		assert.Equalf(t, float64(tt.result), v, messageForFailedTest(i, tt.description))
	}
}

func TestToInt(t *testing.T) {
	cases := []struct {
		input                interface{}
		result               int64
		wantError            bool
		errorMessageContains string
		description          string
	}{
		{
			"Hello",
			0,
			true,
			"invalid syntax",
			"convert string that is not a number",
		},
		{
			struct{ A int }{A: 5},
			0,
			true,
			"Value struct { A int } is type struct",
			"convert struct to int",
		},
		{
			float64(179769313.4862932),
			179769313,
			false,
			"",
			"convert float64",
		},
		{
			float32(359.9),
			359,
			false,
			"",
			"convert float32",
		},
		{
			int64(9223372036854775807),
			9223372036854775807,
			false,
			"",
			"convert int64",
		},
		{
			int32(2147483647),
			2147483647,
			false,
			"",
			"convert int32",
		},
		{
			int16(32767),
			32767,
			false,
			"",
			"convert int16",
		},
		{
			int8(127),
			127,
			false,
			"",
			"convert int8",
		},
		{
			int(6),
			6,
			false,
			"",
			"convert int",
		},
		{
			"1",
			1,
			false,
			"",
			"convert string",
		},
	}
	for i, tt := range cases {
		v, e := iconvert.ToInt(tt.input)
		if tt.wantError {
			assert.Errorf(t, e, messageForFailedTest(i, tt.description))
			if tt.errorMessageContains != "" {
				assert.Containsf(t, e.Error(), tt.errorMessageContains, messageForFailedTest(i, tt.description))
			}

			continue
		}
		assert.NoErrorf(t, e, messageForFailedTest(i, tt.description))
		assert.Equalf(t, int64(tt.result), v, messageForFailedTest(i, tt.description))
	}
}
