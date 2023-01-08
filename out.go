package slice

import (
	"encoding/json"
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

//func AppendCSVFile() {
//
//}

//// SaveAsCSVFile 把矩阵保存为CSV文件
//func SaveAsCSVFile[T any](slice []T, filepath string) error {
//
//	for {
//
//	}
//
//	file, err := os.Open(filepath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	writer := csv.NewWriter(file)
//	writer.Write()
//	return nil
//}

// SaveAsJsonLineFile 把切片保存到JSON文件
func SaveAsJsonLineFile[T any](slice []T, filepath string) error {
	marshal, err := json.Marshal(slice)
	if err != nil {
		return err
	}
	open, err := os.Open(filepath)
	if err != nil {
		return err
	}
	_, err = open.Write(marshal)
	return err
}
