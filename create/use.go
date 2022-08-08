package create

import "fmt"

func UseSql(sql string) {
	dbName := sql[4 : len(sql)-1]

	path = getPath()
	dataBases := GetAllDataBases(path)
	for _, dataBase := range dataBases {
		if dbName == dataBase {
			setNowPath(dbName)
			fmt.Println("Database changed")
			Get()
		} else {
			fmt.Println("ERROR:数据库不存在")
			Get()
		}
	}
}
