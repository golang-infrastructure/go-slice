package slice

import (
	"bufio"
	"os"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// ReadAllLines 读取文件中的所有行
func ReadAllLines(filepath string) ([]string, error) {
	return ReadLines(filepath, 0, -1)
}

// ReadLines 按行读取文件，跳过skip行，读取limit行
func ReadLines(filepath string, skipLine, limitLine int) ([]string, error) {
	open, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(open)

	// skip lines
	for skipLine > 0 {
		scanner.Text()
		skipLine--
	}

	lines := make([]string, 0)
	for scanner.Scan() && limitLine != 0 {
		lines = append(lines, scanner.Text())
		limitLine--
	}
	return lines, nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

//// ReadAllBytes 读取整个文件，以字节数组返回
//func ReadAllBytes(filepath string) ([]byte, error) {
//	return ReadBytes(filepath, 0, -1)
//}

//// ReadBytes 从文件中读取给定数量的字节
//func ReadBytes(filepath string, skipBytes, bytesLimit int) ([]byte, error) {
//	//open, err := os.Open(filepath)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//scanner := bufio.NewScanner(open)
//	//
//	//// skip bytes
//	//if skipBytes > 0 {
//	//	skipBytes--
//	//}
//	//
//	//bytes := make([]byte, 0)
//	//for scanner.Scan() && bytesLimit != 0 {
//	//	bytes = append(bytes, scanner.Bytes()...)
//	//	bytesLimit--
//	//}
//	//return bytes, nil
//	return nil, nil
//}

// ------------------------------------------------ ---------------------------------------------------------------------

//func ReadCSV() [][]string {
//	return nil
//}
//
//func ReadJsonArray(filepath string) []string {
//	return nil
//}
