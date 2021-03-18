package _317_2

import (
	"strconv"
	"strings"
)

func originalDigits(s string) string {
	count := make([]int, 26+'a') // a-z 数量为26，+a为最小值
	for _, b := range s {
		count[b]++
	}

	out := make([]int, 10)
	out[0] = count['z']
	out[2] = count['w']
	out[4] = count['u']
	out[6] = count['x']
	out[8] = count['g']
	out[1] = count['o'] - out[0] - out[2] - out[4]
	out[3] = count['h'] - out[8]
	out[5] = count['f'] - out[4]
	out[7] = count['s'] - out[6]
	out[9] = count['i'] - out[5] - out[6] - out[8]

	//res := ""
	var res strings.Builder // 效率更高
	for i, v := range out {
		for j := 0; j < v; j++ {
			//res += strconv.Itoa(i)
			res.WriteString(strconv.Itoa(i))
		}
	}

	return res.String()
}
