package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//lines := ReadFileByLines("2.txt")
	ReadFileByLine("2.txt")
	//fmt.Println(lines)
}
func CreateFile() {
	file, err := os.OpenFile("2.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer file.Close()
	for i := 0; i < 3; i++ {
		file.Write([]byte("type"))
		file.Write([]byte("\n"))
	}
	file.Write([]byte("python"))
	file.Write([]byte("\n"))
	file.Write([]byte("java"))
	file.Write([]byte("\n"))
	file.Write([]byte("go"))
	file.Write([]byte("\n"))
}
func ReadFileByLines(path string) []string {
	list := []string{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open error")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "***" {
			list = append(list, line)
		} else {
			break
		}
	}
	return list
}
func ReadFileByLine(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open error")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(count)
		fmt.Println(line)
		count++
	}
}
