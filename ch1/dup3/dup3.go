// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.

// ファイルから全行を一括して読み込む
// 練習問題1.4

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile : ファイルの内容全体を読み込む．stringに変換されるであろうバイトのスライスに変換
		data, err := ioutil.ReadFile(filename)
		fmt.Println(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if filenames[line] == nil {
				filenames[line] = make(map[string]int)
			}
			filenames[line][filename]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, n := range filenames[line] {
				fmt.Printf("  %s\t%d\n", filename, n)
			}
		}
	}
}

//!-
