/**
 * @Author: alessonhu
 * @Description:
 * @File:  path
 * @Version: 1.0.0
 * @Date: 2022/5/5 12:26
 */

package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetMainPath(filename string) string {
	return GetMainDiectory() + filename
}

func GetMainDiectory() string {
	path, err := filepath.Abs(os.Args[0])

	if err != nil {
		return ""
	}

	full_path := filepath.Dir(path)

	return PathAddBackslash(full_path)
}

func PathAddBackslash(path string) string {
	i := len(path) - 1

	if !os.IsPathSeparator(path[i]) {
		path += string(os.PathSeparator)
	}

	return path
}

func GetExeFileBaseName() string {
	name := GetExeFileName()
	return strings.TrimSuffix(name, filepath.Ext(name))
}

func GetExeFileName() string {
	path, err := filepath.Abs(os.Args[0])

	if err != nil {
		return ""
	}

	return filepath.Base(path)
}

func PathFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func ReadFileAsString(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}
