package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var m string // mark
var r string // return
var l string // list
var e string // edit
var t string // test
var d string // delete
var c string // clean

func init() {
	flag.StringVar(&m, "m", "", "标记目录 默认标记为当前目录名称")
	flag.StringVar(&r, "r", "", "返回至标记目录")
	flag.StringVar(&l, "l", "", "列出所有标记")
	flag.StringVar(&e, "e", "", "使用vim编辑标记")
	flag.StringVar(&t, "t", "", "测试标记文件是否符合规范以及查看标记目录是否存在")
	flag.StringVar(&d, "d", "", "删除标记")
	flag.StringVar(&c, "c", "", "清除无用标记(标记目录不存在)")
	flag.Parse()

	// todo 命令输入处理
}

func main() {
	// 命令模式 ? 策略模式
	fmt.Println("mark:", m)
	readMR()
}

func readMR() {
	filePath := "/Users/lkzc19/Downloads/.mr"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
