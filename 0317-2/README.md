# [面试 & 力扣423. 从英文中重建数字](https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/)


## 题目

给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。

注意:

输入只包含小写英文字母。
输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。
输入字符串的长度小于 50,000。
示例 1:
```
输入: "owoztneoer"

输出: "012" (zeroonetwo)
```

示例 2:
```
输入: "fviefuro"

输出: "45" (fourfive)
```


## 解题思路

这道题我没有想到，居然是找规律。。。

因为题目规定英文单词表示的数字0-9（英文不好的同学想骂娘了）

0的英文是zero，1的英文是one，2的英文是two，3的英文是three，4的英文是four，5的英文是five，6的英文是six，7的英文是seven，8的英文是eight，9的英文是nine。

因此，我们需要寻找一些独特的标志。注意到，所有的偶数都包含一个独特的字母：

“z” 只在 “zero” 中出现。
“w” 只在 “two” 中出现。
“u” 只在 “four” 中出现。
“x” 只在 “six” 中出现。
“g” 只在 “eight” 中出现。

由已知可以推出剩余未知：

“o” 在 “one”，""zero", “two” 和 “four” 中出现, 可以推出 1
“h” 只在 “three” 和 “eight” 中出现。可以推出 3
“f” 只在 “five” 和 “four” 中出现。可以推出 5
“s” 只在 “seven” 和 “six” 中出现。可以推出 7
“i” 在 “nine”，“five”，“six” 和 “eight” 中出现。可以推出 9


## 代码

```go
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
```


![](http://wesub.ifree258.top/bottomPic.png)