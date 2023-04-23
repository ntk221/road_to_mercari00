package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// func ReadDir(name string) ([]DirEntry, error)
// func Walk(root string, fn WalkFunc) error

// あるディレクトリの配下のファイル一覧を取得する
func listFiles(dirPath string) ([]string, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("error: %s: no such file or directory", dirPath)
	}
	files := []string{}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}
	files := []string{}
	for _, dir := range args {
		dirFiles, err := listFiles(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		files = append(files, dirFiles...)
	}

	// ファイルのパスを出力
	for _, file := range files {
		fmt.Println(file)
	}
}
