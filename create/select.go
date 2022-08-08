package create

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func selectSql(sql string) {
	sep := getSeparate()
	path := getNowPath()
	tableName := sql[14 : len(sql)-1]
	nowPath := path + "/" + tableName + ".txt"
	list := getTable(nowPath)
	restrict1 := sql[7:8]
	b := strings.Contains(tableName, " ")
	if !b && restrict1 == "*" {
		columnName := getColumnName(nowPath, 1)
		columnName = columnName[0 : len(columnName)-1]
		lists := [][]string{}
		for _, value := range list {
			list1 := []string{}
			strings := strings.Split(value, sep)
			strings = strings[0 : len(strings)-1]
			for _, value := range strings {
				list1 = append(list1, value)
			}
			lists = append(lists, list1)
		}
		fmt.Println(generateTable(columnName, lists))
		fmt.Println("Query OK")
	} else {
		fmt.Println("ERROR: 目前的select命令不支持这种方式的查询")
	}
	Get()
}
func getTable(path string) []string {
	list := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", path, err)
	}
	defer file.Close()
	index := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		if index > 4 {
			list = append(list, scanner.Text())
		}
	}
	return list
}
