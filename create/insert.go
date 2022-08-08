package create

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var tableName string

func insertSql(sql string) {
	index := strings.Index(sql, "(")
	name := sql[12:index]
	tableName = strings.Trim(name, " ")

	path := getNowPath()
	list := getAllTables(path)
	var tableExist bool = false
	for _, value := range list {
		if tableName == value {
			tableExist = true
		}
	}
	if tableExist {
		list1 := []string{}
		reg1 := regexp.MustCompile("\\(.*?\\)")
		matcher := reg1.FindAll([]byte(sql), -1)
		for _, v := range matcher {
			list1 = append(list1, string(v))
		}
		if len(list1) != 2 {
			fmt.Println("ERROR: 语句有错误")
			Get()
		} else {
			s1 := strings.Trim(list1[0][1:len(list1[0])-1], "")
			s2 := strings.Trim(list1[1][1:len(list1[1])-1], "")
			s2Transmean := transMean(s2)
			key := strings.Split(s1, ",")
			value := strings.Split(s2Transmean, ",")
			for i := 0; i < len(key); i++ {
				key[i] = strings.Trim(key[i], " ")
			}
			for i := 0; i < len(value); i++ {
				value[i] = strings.Trim(value[i], "\"")
				value[i] = strings.Trim(value[i], "'")
			}
			sep := getSeparate()
			str := ""
			for i := 0; i < len(value); i++ {
				str += value[i] + sep
			}
			writeFile(path+"/"+tableName+".txt", str)
			fmt.Println("Query OK")
			Get()
		}

	} else {
		fmt.Println("ERROR:该表不存在")
		Get()
	}
}
func getColumnName(path string, i int) []string {
	list := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", path, err)
	}
	defer file.Close()
	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		if index == i {
			sep := getSeparate()
			line := scanner.Text()
			split := strings.Split(line, sep)
			for _, value := range split {
				list = append(list, value)
			}
		}
	}
	return list

}
func transMean(str string) string {
	sep := getSeparate()
	all := strings.ReplaceAll(str, sep, sep+sep)
	return all
}
func writeFile(path, s string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666) //0666 在windos下无效
	defer file.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	s += "\n"
	file.Write([]byte(s))
}
func check(strType, str string) string {
	b := strings.Contains(strType, "char")
	if b {
		reg1 := regexp.MustCompile("\".*\"")
		matcher := reg1.FindAll([]byte(str), -1)
		if matcher != nil {
			reg2 := regexp.MustCompile("\"")
			s := reg2.ReplaceAll([]byte(str), []byte(""))
			return string(s)
		} else {
			fmt.Println("ERROR: 输入值的格式错误")
		}
	} else {
		return "lianhonghui"
	}
	return ""
}
