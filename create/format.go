package create

import (
	"regexp"
	"strings"
)

func FormatSQL(str string) string {
	strTrim := strings.Trim(str, " ")
	strLower := strings.ToLower(strTrim)
	reg := regexp.MustCompile("\\s{2,}")
	strReplaceSpace := reg.ReplaceAll([]byte(strLower), []byte(" "))
	reg1 := regexp.MustCompile("( ;)$")
	strReplace := reg1.ReplaceAll(strReplaceSpace, []byte(";"))
	return string(strReplace)
}
