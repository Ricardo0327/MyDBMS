package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//sql := "create table (name varchar(10) not null,age int not null)"
	sql1 := "insert into one (name,age) values ('machi',10);"
	reg1, _ := regexp.Compile("\\(.*?\\)")
	matcher := reg1.FindString(sql1)
	matcher = matcher[1 : len(matcher)-1]
	split := strings.Split(matcher, ",")
	fmt.Println(split)
}
