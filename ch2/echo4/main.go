// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// n, sepはフラグ変数へのポインタ
// flag.Bool bool型のフラグ変数を生成 -> アクセスは*n
var n = flag.Bool("n", false, "omit trailing newline")

// flag.String string型のフラグ変数を生成 -> アクセスは*sep
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	// flag.Args フラグでない引数を取得する
	fmt.Println(flag.Args())
	fmt.Println(*sep)
	fmt.Println(*n)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//!-
