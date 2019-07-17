package util

import (
	"os"
	"path/filepath"
	"strings"
)

// 当前程序执行所在的目录
func CurrentDir() string {
	current, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return current
}

// 所有_test.go结尾的文件
func GoTestFiles(root string) []string {
	slc := []string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), "_test.go") {
			slc = append(slc, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return slc
}
