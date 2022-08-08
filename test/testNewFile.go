package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.OpenFile("D:\\go\\mszlu\\MyDatabase\\DataBase\\1.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {

		fmt.Printf("open file failed,err:%v", err)
		return

	}
}
