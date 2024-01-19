package decode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeBencode(t *testing.T) {
	// Test case for decoding a string
	t.Run("Decode String", func(t *testing.T) {
		input := "5:hello"
		expected := "hello"

		result, err := decodeBencode(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
			fmt.Printf("Decoded result: %v\n", result)
			fmt.Printf("Error: %v\n", err)
		}
	})

	// Test case for decoding an integer
	t.Run("Decode Integer", func(t *testing.T) {
		input := "i42e"
		expected := 42

		result, err := decodeBencode(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
			fmt.Printf("Decoded result: %v\n", result)
			fmt.Printf("Error: %v\n", err)
		}
	})

	// Test case for decoding a list
	t.Run("Decode List", func(t *testing.T) {
		input := "l3:one3:twoe"
		expected := []interface{}{"one", "two"}

		result, err := decodeBencode(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
			fmt.Printf("Decoded result: %v\n", result)
			fmt.Printf("Error: %v\n", err)
		}
	})

	// Test case for decoding an empty dictionary
	t.Run("Decode Empty Dictionary", func(t *testing.T) {
		input := "de"
		expected := map[string]interface{}{}

		result, err := decodeDictionary([]byte(input))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// Test case for decoding a dictionary with string values
	t.Run("Decode Dictionary with Strings", func(t *testing.T) {
		input := "d3:key53:value3:foo3:bare"
		expected := map[string]interface{}{"key": "value", "foo": "bar"}

		result, err := decodeDictionary([]byte(input))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
