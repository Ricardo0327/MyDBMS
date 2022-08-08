package main

import (
	"fmt"
	"math"
	"path"
	"runtime"
)

func main() {
	//fmt.Println("getCurrentAbPathByCaller = ", getCurrentAbPathByCaller())
	fmt.Println(math.MinInt)
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
