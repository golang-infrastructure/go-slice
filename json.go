package slice

import "encoding/json"

func JsonUnmarshal[T any](slice []byte, value T) error {
	return json.Unmarshal(slice, value)
}
