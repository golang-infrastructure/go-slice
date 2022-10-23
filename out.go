package slice

import (
	"os"
	"strings"
)

// SaveAsTextFile 保存到文本文件
func SaveAsTextFile[T any](slice []T, filepath string) error {
	stringSlice := ToStringSlice(slice)
	open, err := os.Open(filepath)
	if err != nil {
		return err
	}
	_, err = open.WriteString(strings.Join(stringSlice, "\n"))
	return err
}

// SaveAsCSVFile 保存到CSV文件
func SaveAsCSVFile[T any](slice []T, filepath string) error {
	return nil
}

// SaveAsJsonFile 保存到JSON文件
func SaveAsJsonFile[T any](slice []T, filepath string) error {
	return nil
}
