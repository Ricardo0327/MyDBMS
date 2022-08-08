package main

import (
	"fmt"
	"regexp"
)

func main() {
	findTest()
	//findIndexTest()
	//findStringTest()
	//findChinesString()
	//findNumOrLowerLetter()
	//findAndReplace()
}

//传入[]byte，返回[]byte
func findTest() {
	str := "ab001234hah120210a880218end"
	reg := regexp.MustCompile("\\d{6}") //六位连续的数字
	fmt.Println("------Find------")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	fmt.Println(string(data))
	fmt.Println("------FindAll------")
	//返回str中所有匹配reg的字符串
	//第二个参数表示最多返回的个数，传-1表示返回所有结果
	dataSlice := reg.FindAll([]byte(str), -1)
	for _, v := range dataSlice {
		fmt.Println(string(v))
	}
	fmt.Println("------FindSQL------")
	sql1 := "insert into one (name,age) values ('machi',10);"
	reg1 := regexp.MustCompile("\\(.*?\\)")
	matcher := reg1.FindAll([]byte(sql1), -1)
	for _, v := range matcher {
		fmt.Println(string(v))
	}
}

//传入[]byte，返回首末位置索引
func findIndexTest() {
	fmt.Println("------FindIndex------")
	//返回第一个匹配的字符串的首末位置
	reg2 := regexp.MustCompile("start\\d*end") //start开始，end结束，中间全是数字
	str2 := "00start123endhahastart120PSend09start10000end"
	//index[0]表示开始位置，index[1]表示结束位置
	index := reg2.FindIndex([]byte(str2))
	fmt.Println("start:", index[0], ",end:", index[1], str2[index[0]:index[1]])
	fmt.Println("------FindAllIndex------")
	//返回所有匹配的字符串首末位置
	indexSlice := reg2.FindAllIndex([]byte(str2), -1)
	for _, v := range indexSlice {
		fmt.Println("start:", v[0], ",end:", v[1], str2[v[0]:v[1]])
	}
}

//传入string，返回string(更加方便)
func findStringTest() {
	fmt.Println("------FindString------")
	str := "ab001234hah120210a880218end"
	reg := regexp.MustCompile("\\d{6}") //六位连续的数字
	fmt.Println(reg.FindString(str))
	fmt.Println(reg.FindAllString(str, -1))
	//以下两个方法是类似的
	fmt.Println(reg.FindStringIndex(str))
	fmt.Println(reg.FindIndex([]byte(str)))
}

//查找汉字
func findChinesString() {
	str := "hello中国hello世界和平hi好"
	reg := regexp.MustCompile("[\\p{Han}]+")
	fmt.Println(reg.FindAllString(str, -1))
	//[中国 世界和平 好]
}

//查找数字或小写字母
func findNumOrLowerLetter() {
	str := "HAHA00azBAPabc09FGabHY99"
	reg := regexp.MustCompile("[\\d|a-z]+")
	fmt.Println(reg.FindAllString(str, -1))
	//[00az abc09 ab 99]
}

//查找并替换
func findAndReplace() {
	str := "Welcome for Beijing-Tianjin CRH train."
	reg := regexp.MustCompile(" ")
	fmt.Println(reg.ReplaceAllString(str, "@")) //将空格替换为@字符
	//Welcome@for@Beijing-Tianjin@CRH@train.
}
