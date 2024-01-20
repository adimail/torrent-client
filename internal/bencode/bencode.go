package bencode

import (
	"errors"
	"fmt"
	"strconv"
)

func DecodeBencode(data string) (interface{}, error) {
	return decodeBencodeHelper([]byte(data))
}

func decodeBencodeHelper(data []byte) (interface{}, error) {
	switch data[0] {
	case 'i':
		return decodeInteger(data[1:])
	case 'l':
		return decodeList(data[1:])
	case 'd':
		return decodeDictionary(data[1:])
	default:
		return decodeString(data)
	}
}

func decodeInteger(data []byte) (int, error) {
	endIndex := indexOf(data, 'e')
	if endIndex == -1 {
		return 0, errors.New("invalid bencode format: missing 'e' for integer")
	}

	intValue, err := strconv.Atoi(string(data[:endIndex]))
	if err != nil {
		return 0, errors.New("Invalid Bencode format: Unable to parse integer")
	}

	return intValue, nil
}

func decodeString(data []byte) (string, error) {
	colonIndex := indexOf(data, ':')
	if colonIndex == -1 {
		return "", errors.New("Invalid Bencode format: Missing ':' for string length")
	}

	length, err := strconv.Atoi(string(data[:colonIndex]))
	if err != nil {
		return "", errors.New("Invalid Bencode format: Unable to parse string length")
	}

	startIndex := colonIndex + 1
	endIndex := startIndex + length
	if endIndex > len(data) {
		return "", errors.New("Invalid Bencode format: String length exceeds available data")
	}

	return string(data[startIndex:endIndex]), nil
}

func decodeList(data []byte) ([]interface{}, error) {
	var list []interface{}

	for len(data) > 0 && data[0] != 'e' {
		item, err := decodeBencodeHelper(data)
		if err != nil {
			return nil, err
		}

		list = append(list, item)

		strLen := len(fmt.Sprintf("%v", item))
		data = data[strLen+len(strconv.Itoa(strLen))+1:]
	}

	return list, nil
}

func decodeDictionary(data []byte) (map[string]interface{}, error) {
	dict := make(map[string]interface{})

	for len(data) > 0 && data[0] != 'e' {
		key, err := decodeString(data)
		if err != nil {
			return nil, err
		}

		keyLen := len(fmt.Sprintf("%v", key))
		data = data[keyLen+len(strconv.Itoa(keyLen))+1:]

		value, err := decodeBencodeHelper(data)
		if err != nil {
			return nil, err
		}

		valLen := len(fmt.Sprintf("%v", value))
		data = data[valLen+len(strconv.Itoa(valLen))+1:]

		dict[key] = value
	}

	return dict, nil
}

func indexOf(data []byte, target byte) int {
	for i, b := range data {
		if b == target {
			return i
		}
	}
	return -1
}
