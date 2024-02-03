package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
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

var mrPath string

var mr = make(map[string]string)

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

	read()

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
	cmd := exec.Command("cd", "/Users/lkzc19/Projects/lkzc19/MarkReturn")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("无法切换目录:", err)
		return
	}

	//list()
	//mark()
}

func mark() {
	currentDir, err := os.Getwd()
	check(err)

	var key string
	if m != "" {
		key = m
	} else {
		key = filepath.Base(currentDir)
	}

	value, ok := mr[key]

	if ok {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("是否要覆盖[key=" + value + "]? (Y/n): ")

		input, err := reader.ReadString('\n')
		check(err)
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "n" {
			return
		} else {
			mr[key] = currentDir
		}
	} else {
		mr[key] = currentDir
	}

	write()
}

// ret 即 return
func ret() {

}

func list() {
	for key, value := range mr {
		fmt.Println(key + "\t->\t" + value)
	}
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

func read() {
	mrPath = getFilePath()
	_, err := os.Stat(mrPath)
	if os.IsNotExist(err) {
		_, err := os.Create(mrPath)
		check(err)
	} else if err == nil {
		content, err := os.ReadFile(getFilePath())
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), "\n")
		for _, word := range lines {
			if strings.TrimSpace(word) == "" {
				continue
			}
			tmp := strings.Split(strings.TrimSpace(word), "=")
			if len(tmp) != 2 {
				fmt.Println(".mr 文件错误")
			}
			mr[tmp[0]] = tmp[1]
		}
	} else {
		panic(err)
	}
}

func write() {
	var builder strings.Builder
	for key, value := range mr {
		builder.WriteString(key + "=" + value + "\n")
	}

	// 打开文件，以写入的方式
	file, err := os.OpenFile(mrPath, os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	// 写入数据
	_, err = file.Write([]byte(builder.String()))
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
