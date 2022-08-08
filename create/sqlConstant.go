package create

import "fmt"

var path string = "../MyDatabase/DataBase"
var nowPath string = path
var separate string = "~"

func getPath() string {
	return path
}
func getNowPath() string {
	return nowPath
}
func setNowPath(name string) {
	nowPath = nowPath + "/" + name
	fmt.Println(nowPath)
}
func getSeparate() string {
	return separate
}
