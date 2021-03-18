# [面试 offer]


## 题目

一头母羊的寿命是5年,它会在第2年底和第4年底各生下一头母羊,第5年底死去,
问一开始农场有1头母羊,N年后,农场会有多少只母羊？


## 解题思路

这种"一生二、二生三、三生万物"的问题，就想着递归算法（顺便和斐波那契数列联系起来）

代码实现很简单，难点在于你是否能转过来弯。

## 代码

```go
func getSheep(year int) int {
	num := 1 // init sheep num
	// note: init i = 1; i <= year
	for i := 1; i <= year; i++ {
		if i == 2 {
			num += getSheep(year-2) // 递归计算2年时出生的新羊
		}else if i == 4 {
			num += getSheep(year-4) // 递归计算4年时出生的新羊
		}else if i == 5 {
			num--
			break // 优化
		}
	}
	return num
}
```


![](http://wesub.ifree258.top/bottomPic.png)