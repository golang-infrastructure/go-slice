package slice

import "encoding/json"

func BytesJsonUnmarshal[T any](slice []byte, value T) error {
	return json.Unmarshal(slice, value)
}
