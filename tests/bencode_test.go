package bencode_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/adimail/torrent-client/internal/bencode"
)

type testcases struct {
	input   string
	expect  interface{}
	comment string
}

func TestDecodeBencode(t *testing.T) {
	// string
	t.Run("Decode String", func(t *testing.T) {
		testCases := []testcases{
			{"11:hello World", "hello World", "Decode string"},
			{"0:", "", "Decode empty string"},
			{"7:1234567", "1234567", "Decode string with digits"},
		}

		for _, tc := range testCases {
			t.Run(tc.comment, func(t *testing.T) {
				result, err := bencode.DecodeBencode(tc.input)

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tc.expect) {
					t.Errorf("Expected %v, but got %v", tc.expect, result)
					fmt.Printf("Decoded result: %v\n", result)
					fmt.Printf("Error: %v\n", err)
				}
			})
		}
	})

	// Test case for decoding an integer
	t.Run("Decode Integer", func(t *testing.T) {
		testCases := []testcases{
			{"i15e", 15, "Decode positive integer"},
			{"i-42e", -42, "Decode negative integer"},
		}

		for _, tc := range testCases {
			t.Run(tc.comment, func(t *testing.T) {
				result, err := bencode.DecodeBencode(tc.input)

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tc.expect) {
					t.Errorf("Expected %v, but got %v", tc.expect, result)
					fmt.Printf("Decoded result: %v\n", result)
					fmt.Printf("Error: %v\n", err)
				}
			})
		}
	})

	// list
	t.Run("Decode List", func(t *testing.T) {
		testCases := []testcases{
			{"l3:one3:twoe", []interface{}{"one", "two"}, "Decode list of strings"},
		}

		for _, tc := range testCases {
			t.Run(tc.comment, func(t *testing.T) {
				result, err := bencode.DecodeBencode(tc.input)

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tc.expect) {
					t.Errorf("Expected %v, but got %v", tc.expect, result)
					fmt.Printf("Decoded result: %v\n", result)
					fmt.Printf("Error: %v\n", err)
				}
			})
		}
	})

	// Dictionary
	t.Run("Decode Dictionary", func(t *testing.T) {
		testCases := []testcases{
			{"de", map[string]interface{}{}, "Decode empty dictionary"},
			{"d3:key5:value3:foo3:bare", map[string]interface{}{"key": "value", "foo": "bar"}, "Decode dictionary with strings"},
			{"d3:cow3:moo3:dog4:barke", map[string]interface{}{"cow": "moo", "dog": "bark"}, "Decode dictionary with strings"},
		}

		for _, tc := range testCases {
			t.Run(tc.comment, func(t *testing.T) {
				result, err := bencode.DecodeBencode(tc.input)

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(result, tc.expect) {
					t.Errorf("Expected %v, but got %v", tc.expect, result)
					fmt.Printf("Decoded result: %v\n", result)
					fmt.Printf("Error: %v\n", err)
				}
			})
		}
	})
}
