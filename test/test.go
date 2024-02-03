package main

import (
	"fmt"
	"log"
	"os"
)

var directories = map[string]string{
	"docs":      "/Users/lkzc19/Projects/lkzc19",
	"projects":  "/path/to/projects",
	"downloads": "/path/to/downloads",
}

func gotoDirectory(dirAlias string) {
	dirPath, found := directories[dirAlias]
	if !found {
		fmt.Printf("未找到目录别名 '%s'\n", dirAlias)
		return
	}

	err := os.Chdir(dirPath)
	if err != nil {
		log.Fatalf("无法切换到目录 '%s': %s\n", dirAlias, err)
	}

	fmt.Printf("切换到目录 '%s' (%s)\n", dirAlias, dirPath)
}

func printDirectoryAliases() {
	fmt.Println("可用的目录别名：")
	for dirAlias, dirPath := range directories {
		fmt.Printf("- %s: %s\n", dirAlias, dirPath)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printDirectoryAliases()
	} else {
		gotoDirectory(args[0])
	}
}
