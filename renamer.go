package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func rename(name string) string {
	tmp := strings.ToLower(string(name[0]))
	for _, j := range name[1:] {
		if unicode.IsUpper(j) {
			tmp += "-" + string(unicode.ToLower(j))
			continue
		}
		tmp += string(j)
	}
	return tmp
}

func main() {
	var root string
	flag.StringVar(&root, "path", "", "where you script will be loaded")
	flag.Parse()
	absPath, _ := os.Getwd()
	root = path.Join(absPath, root)
	dirs, _ := os.ReadDir(root)
	fmt.Println(root)
	for _, dir := range dirs {
		//os.Rename(dir.Name(), rename(dir.Name()))
		os.Rename(path.Join(root, dir.Name()), path.Join(root, rename(dir.Name())))
		//fmt.Println(path.Join(root, dir.Name()), path.Join(root, rename(dir.Name())))
		//break
	}
}
