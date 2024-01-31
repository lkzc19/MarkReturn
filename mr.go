package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

var version = "0.0.1"

var (
	m string // mark
	r string // ret
	l bool   // list
	e bool   // edit
	t bool   // test
	d string // del
	c bool   // clean
)

var path string

var file *os.File

func init() {
	flag.Usage = usage
	flag.StringVar(&m, "m", "", "标记目录 默认标记为当前目录名称")
	flag.StringVar(&r, "r", "", "返回至标记目录")
	flag.BoolVar(&l, "l", false, "列出所有标记")
	flag.BoolVar(&e, "e", false, "使用vim编辑标记")
	flag.BoolVar(&t, "t", false, "测试标记文件是否符合规范以及查看标记目录是否存在")
	flag.StringVar(&d, "d", "", "删除标记")
	flag.BoolVar(&c, "c", false, "清除无用标记(标记目录不存在)")
	flag.Parse()

	// todo 命令输入处理
	path = getFilePath()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

	} else if err == nil {

	} else {
		panic(err)
	}

	content, err := os.ReadFile(getFilePath())
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(content), "\n")
	for _, word := range words {
		split := strings.Split(strings.TrimSpace(word), "=")
		fmt.Println(split)

	}
	//file, err := os.OpenFile(getFilePath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//mr := make(map[string]string)

	//r := bufio.NewReader(file)
	//for {
	//	line, _, err := r.ReadLine()
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		panic(err)
	//	}
	//	s := strings.TrimSpace(string(line))
	//	index := strings.Index(s, "=")
	//	if index < 0 {
	//		continue
	//	}
	//	key := strings.TrimSpace(s[:index])
	//	if len(key) == 0 {
	//		continue
	//	}
	//	value := strings.TrimSpace(s[index+1:])
	//	if len(value) == 0 {
	//		continue
	//	}
	//	mr[key] = value
	//}
	//fmt.Println(mr)

	// 写入数据到文件
	//_, err = file.Write(data)
	//if err != nil {
	//	fmt.Println("无法写入文件:", err)
	//	return
	//}
}

func usage() {
	// 只是为了美观
	versionText := "MarkReturn version: mr/" + version + "\n"
	// todo 修改用法提示
	usageText := "Usage: mr [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]\n\n"
	optionsText := "Options:\n"
	_, err := fmt.Fprintf(os.Stderr, versionText+usageText+optionsText)
	check(err)
	flag.PrintDefaults()
}

func main() {
	//fmt.Println("list:", l)
	//readMR()
}

func mark() {

}

// ret 即 return
func ret() {

}

func list() {

}

func edit() {

}

func test() {

}

// del 即 delete
func del() {

}

func clean() {

}

func getFilePath() string {
	tmp := os.Getenv("MR")
	if tmp == "" {
		currentUser, err := user.Current()
		if err != nil {
			panic(err)
		}
		return currentUser.HomeDir + "/.mr"
	} else {
		return tmp + "/.mr"
	}
}

func checkFile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

	} else if err == nil {

	} else {
		panic(err)
	}
}

//func read() map[string]string {
//	_, err := os.Stat(path)
//	if os.IsNotExist(err) {
//		return map[string]string{}
//	} else if err == nil {
//		data, err := os.ReadFile(path)
//		check(err)
//
//	} else {
//		panic(err)
//	}
//}

func write() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
