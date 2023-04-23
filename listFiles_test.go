package main

import (
	"testing"
	"os"
	"path/filepath"
)

// func MkdirAll(path string, perm FileMode) error
// func WriteFile(name string, data []byte, perm FileMode) error

func TestListFiles(t *testing.T) {
	dir := "/tmp/testdir"
	files := []string{
		dir + "/file1.txt",
	}
	for _, file := range files {
		_ = os.MkdirAll(filepath.Dir(file), 0777)	
		_ = os.WriteFile(file, []byte(""), 0666)
	}
	defer os.RemoveAll(dir)

	// テストケース
	testCases := []struct {
		name		string
		dirPath		string
		expected	[]string
		expectErrors	bool
	}{
		{
			name: "all files in test directory",
			dirPath: dir,
			expected: []string{
				"/tmp/testdir/file1.txt",
			},
			expectErrors: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			files, err := listFiles(tc.dirPath)
			if tc.expectErrors {
				if err == nil {
					t.Error("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(files) != len(tc.expected) {
					t.Errorf("expected %d files but got %d files", len(tc.expected), len(files))
				} else {
					for i, file := range tc.expected {
						if file != files[i] {
							t.Errorf("expected %s but got %s", file, files[i])
						}
					}
				}
			}
		})
	}
}
