package create

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var tName string

func CreateSql(sql string) {
	if strings.Contains(sql, "database") == true {
		dbName := sql[16 : len(sql)-1]
		fmt.Println(dbName)
		createDir(dbName)
	} else if strings.Contains(sql, "table") == true {
		createTable(sql)
	}
}
func createDir(dbName string) {
	//basePath是固定目录路径
	basePath := getPath()
	folderPath := filepath.Join(basePath, dbName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
		fmt.Println("Query OK")
		Get()
	} else {
		fmt.Println("ERROR: database exists")
		Get()
	}

}
func createTable(sql string) {
	tablePath := getNowPath()
	index := strings.Index(sql, "(")
	if index != -1 {
		tName = sql[13:index] + ".txt"
		tName = strings.Trim(tName, " ")
		filePath := tablePath + "\\" + tName
		_, err := os.Stat(filePath)
		if err != nil {
			_, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				fmt.Printf("open file failed,err:%v", err)
				return
			}
			sep := getSeparate()
			reg, _ := regexp.Compile(sep + "")
			reg.ReplaceAll([]byte(sql), []byte(sep+sep))

			reg1, _ := regexp.Compile("\\(.*\\)")
			matcher := reg1.FindString(sql)
			matcher = matcher[1 : len(matcher)-1]
			split := strings.Split(matcher, ",")
			list1 := []string{}
			list2 := []string{}
			list3 := []string{}
			for _, value := range split {
				value = strings.Trim(value, " ")
				splitWords := strings.Split(value, " ")
				list1 = append(list1, splitWords[0])
				list2 = append(list2, splitWords[1])
				for i := 2; i < len(splitWords); i++ {
					if splitWords[i] == "not" && splitWords[i+1] == "null" {
						list3 = append(list3, "not null")
						i++
					} else {
						list3 = append(list3, splitWords[i])
					}
				}
			}
			writeFileList(list1, filePath)
			writeFileList(list2, filePath)
			writeFileList(list3, filePath)
			fmt.Println("Query OK")
		} else {
			fmt.Println("ERROR: 该表已经存在")
		}
	} else {
		fmt.Println("ERROR: 语句有错误")
	}
	Get()
}
func writeFileList(list []string, filePath string) {
	str := ""
	sep := getSeparate()
	for _, value := range list {
		str += value + sep
	}
	str += "\r\n"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	file.Write([]byte(str))
}
func Get() {
	input := ""
	for {
		fmt.Print(">>")
		reader := bufio.NewReader(os.Stdin)
		res, _, _ := reader.ReadLine()
		input += string(res)
		input += " "
		if strings.HasSuffix(string(res), ";") == true {
			break
		}

	}
	sql := FormatSQL(input)
	fmt.Println(sql)
	Analysis(sql)
}
