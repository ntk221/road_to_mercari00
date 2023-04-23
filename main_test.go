package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// テスト用ディレクトリを作成
	testDir := "testdir"
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	// テスト用ファイルを作成
	testFiles := []string{"testfile1.jpg", "testfile2.png"}
	for _, filename := range testFiles {
		file, err := os.Create(testDir + "/" + filename)
		if err != nil {
			t.Fatal(err)
		}
		file.Close()
	}

	// 実行可能ファイルを実行
	cmd := exec.Command("./road_to_mercari", testDir)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	output := stdout.String()
	if !strings.Contains(output, "testdir/testfile1.txt") || !strings.Contains(output, "testdir/testfile2.txt") || !strings.Contains(output, "testdir/testfile3.txt") {
		t.Fatalf("Unexpected output: %s", output)
	}
}
