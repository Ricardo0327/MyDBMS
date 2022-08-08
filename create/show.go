package create

import (
	"fmt"
	"strings"
)

func ShowSql(sql string) {
	endsWith := strings.HasSuffix(sql, " databases;")
	if endsWith == true {
		path := getPath()
		dbList := GetAllDataBases(path)
		db := []string{}
		db = append(db, "Database")
		list := [][]string{}
		for _, value := range dbList {
			ls := []string{}
			ls = append(ls, value)
			list = append(list, ls)
		}
		fmt.Println(generateTable(db, list))
		Get()
	} else {
		endsWith = strings.HasSuffix(sql, " tables;")
		nowPath := getNowPath()
		tableList := getAllTables(nowPath)
		list := []string{}

		index := strings.LastIndex(nowPath, "/")
		dbName := nowPath[index+1 : len(nowPath)]
		list = append(list, dbName)
		tlist := [][]string{}
		for _, t := range tableList {
			list1 := []string{}
			list1 = append(list1, t)
			tlist = append(tlist, list1)
		}
		fmt.Println(generateTable(list, tlist))
		Get()
	}
}
