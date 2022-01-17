package main

import (
	"log"
	"os"

	// 使用 "_" 是对包做初始化操作，但是并不使用包里的标识符，目的是调用对应包内的所有代码文件里定义的 init 函数
	// 即：调用matchers包中的rss.go代码文件里的init函数，注册RSS匹配器
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
