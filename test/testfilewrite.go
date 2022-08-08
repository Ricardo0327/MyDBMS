package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个文件并写入内容
	filePath := "1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) //0666 在windos下无效
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	//关闭文件
	defer file.Close()
	//写入的内容
	str := "你好，小王\n"
	//带缓存写入文件
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//内容是写入到缓存中的，需要Flush()
	writer.Flush()
}
