package create

import (
	"io/ioutil"
	"strings"
)

func GetAllDataBases(nowPath string) []string {
	fileInfo, err := ioutil.ReadDir(nowPath)
	if err != nil {
		panic(err)
	}
	dirs := []string{}
	for _, fi := range fileInfo {
		dirs = append(dirs, fi.Name())
	}
	return dirs
}
func getAllTables(nowPath string) []string {
	list := []string{}
	fileInfo, err := ioutil.ReadDir(nowPath)
	if err != nil {
		panic(err)
	}
	// 遍历这个文件夹
	for _, fi := range fileInfo {
		name := fi.Name()
		index := strings.LastIndex(name, ".")
		tableName := name[0:index]
		list = append(list, tableName)
	}
	return list
}
