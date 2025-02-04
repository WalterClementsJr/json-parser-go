package jsonparser

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	type expect struct {
		key   string
		value any
	}
	testCases := []struct {
		input          string
		expectedOutput any
	}{
		{
			`{"id": null, "name": "Jason", "school": { "name": "Bekerly" }, "age": -10, "phone": [], "dob": "2020" }`,
			map[string]any{"id": nil, "name": "Jason", "age": int64(-10), "dob": "2020", "phone": make([]any, 0), "school": map[string]any{"name": "Bekerly"}},
		},
		{
			`[null, 10, "abc", null, [1, "a", null], {"a": "b", "c": [1, "a"]}]`,
			[]any{nil, int64(10), "abc", nil, []any{int64(1), "a", nil}, map[string]any{"a": "b", "c": []any{int64(1), "a"}}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Type casting error", r)
				}
			}()

			actual := Parse(Tokenize(tt.input))

			if !reflect.DeepEqual(actual, tt.expectedOutput) {
				t.Error("Wrong value parsed, expected", tt.expectedOutput, "got", actual)
			}
		})
	}
}
