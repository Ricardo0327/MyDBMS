package create

import (
	"fmt"
	"regexp"
)

const (
	Create   string = "create"
	Help     string = "help"
	Show     string = "show"
	Use      string = "use"
	Quit     string = "quit"
	Describe string = "describe"
	Insert   string = "insert"
	Select   string = "select"
	Drop     string = "drop"
)

func Analysis(sql string) {
	myRegex, _ := regexp.Compile("^([a-z])+")
	start := myRegex.FindString(sql)
	switch start {
	case Create:
		CreateSql(sql)
	case Help:
		help()
	case Show:
		ShowSql(sql)
	case Use:
		UseSql(sql)
	case Quit:
		quitSql()
	case Describe:
		describeSql(sql)
	case Insert:
		insertSql(sql)
	case Select:
		selectSql(sql)
	default:
		fmt.Println("输入的命令无法识别,可以输入help查看目前支持的sql语句")
		Get()
	}

}
