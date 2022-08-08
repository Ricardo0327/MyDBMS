package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//str := "12345"
	//file, _ := os.OpenFile("D:\\go\\mszlu\\MyDatabase\\1.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	//file.Write([]byte(str))
	//_, err := os.Stat("D:\\go\\mszlu\\MyDatabase\\1.txt")
	//if err != nil {
	//	fmt.Println("该文件不存在")
	//} else {
	//	fmt.Println("该文件存在")
	//}
	findFile("./create")
}
func findFile(dir string) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {
		fmt.Println(fi.Name())
		//// 重复输出制表符，模拟层级结构
		//print(strings.Repeat("\t", num))
		//
		//// 判断是不是目录
		//if fi.IsDir() {
		//	println(`目录：`, fi.Name())
		//	findDir(dir+`/`+fi.Name(), num+1)
		//} else {
		//	println(`文件：`, fi.Name())
		//}
	}
}
