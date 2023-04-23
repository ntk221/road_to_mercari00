package main

import (
	"fmt"
	"os"

	"github.com/ntk221/road_to_mercari/fileutil"
)

// func ReadDir(name string) ([]DirEntry, error)
// func Walk(root string, fn WalkFunc) error

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}
	files := []string{}
	for _, dir := range args {
		dirFiles, err := fileutil.ListFiles(dir)
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
