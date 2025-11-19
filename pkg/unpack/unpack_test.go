package unpack_test

import (
	"testing"

	"github.com/andreyxaxa/String-Unpuck/pkg/unpack"
)

func TestString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{
		// 1
		{
			"a4bc2d5e",
			"aaaabccddddde",
			false,
		},
		// 2
		{
			"abcd",
			"abcd",
			false,
		},
		// 3
		{
			"45",
			"",
			true,
		},
		// 4
		{
			"",
			"",
			false,
		},
		// 5
		{
			`qwe\4\5`,
			"qwe45",
			false,
		},
		// 6
		{
			`qwe\45`,
			"qwe44444",
			false,
		},
	}

	for _, tc := range testCases {
		act, err := unpack.String(tc.input)
		if (err != nil) != tc.err {
			t.Errorf("unpackString(%q) error: %v, expected error: %v", tc.input, err, tc.err)
			continue
		}
		if act != tc.expected {
			t.Errorf("unpackString(%q) = %q, want %q", tc.input, act, tc.expected)
		}
	}
}
