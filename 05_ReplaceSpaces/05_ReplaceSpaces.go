package _5_ReplaceSpaces

import (
	"strings"
)

//import "strings"

func replaceSpace(s string) string {
	//res := strings.ReplaceAll(s, " ", "%20")
	//return res

	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		}else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}