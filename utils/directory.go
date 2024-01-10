package utils

import (
	"errors"
	"os"
)

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	// 使用os.Stat()函数获取指定路径的文件信息（如文件状态、大小等）。
	file, err := os.Stat(path)

	if err == nil {
		if file.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
