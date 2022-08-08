package main

import (
	"MyDatabase/create"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Get()

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
	sql := create.FormatSQL(input)
	fmt.Println(sql)
	create.Analysis(sql)
}
