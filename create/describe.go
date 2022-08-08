package create

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func describeSql(sql string) {
	sep := getSeparate()
	path := getNowPath()
	tableName := sql[9 : len(sql)-1]

	nowPath := path + "\\" + tableName + ".txt"
	list := getTableDescribe(nowPath)
	lists := [][]string{}
	s1 := strings.Split(list[0], sep)
	s2 := strings.Split(list[1], sep)
	s3 := strings.Split(list[2], sep)

	for i := 0; i < len(s1); i++ {
		list1 := []string{}
		list1 = append(list1, s1[i])
		list1 = append(list1, s2[i])
		list1 = append(list1, s3[i])
		lists = append(lists, list1)
	}
	list4 := []string{"Field", "Type", "Extra"}
	fmt.Println(generateTable(list4, lists))
}
func getTableDescribe(path string) []string {
	list := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", path, err)
	}
	defer file.Close()
	index := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if index < 4 {
			line := scanner.Text()
			list = append(list, line)
			index++
		}
	}
	return list

}
