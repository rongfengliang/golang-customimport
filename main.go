package main

import (
	"log"
	// 自定义的地址，实际是从github 导入包
	"rongdemo.com"
)

func main() {
	log.Println(shortid.Generate())
}
