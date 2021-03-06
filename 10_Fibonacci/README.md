# [剑指 Offer 10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

简单｜递归、动规

## 题目

写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
```
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
```
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
```
输入：n = 2
输出：1
```

示例 2：
```
输入：n = 5
输出：5
```

提示：

`0 <= n <= 100`


## 解题思路





## 代码

```go

func fib(n int) int {
	// 递归 time limit
	//if n == 1 || n == 2 {
	//	return n
	//}
	//return fib(n-1) + fib(n-2)

	// dp pass
	//if n < 2 {
	//    return n
	//}
	//dp := make([]int, n+1)
	//dp[0] = 0
	//dp[1] = 1
	//for i := 2; i <= n; i++ {
	//	dp[i] = dp[i-1] + dp[i-2]
	//	dp[i] %= 1000000007
	//}
	//return dp[n]

	// 状压dp pass
	//if n < 2 {
	//	return n
	//}
	//one, two := 0, 1
	//for i := 2; i <= n; i++ {
	//	two = one + two
	//	one = two - one
	//	two = two % 1000000007
	//}
	//return two

	//if n < 2 {
	//	return n
	//}
	//one, two, sum := 0, 1, 0
	//for i := 2; i <= n; i++ {
	//	sum = (one + two) % 1000000007
	//	one = two
	//	two = sum
	//}
	//return sum

	// 循环求余 注意小细节： return one
	// one, two, sum := 0, 1, 0
	// for i := 0; i < n; i++ {
	//     sum = (one+two)%1000000007
	//     one = two
	//     two = sum
	// }
	// return one

	// by the way: the best way :D
	f := []int{0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368,75025,121393,196418,317811,514229,832040,1346269,2178309,3524578,5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,267914296,433494437,701408733,134903163,836311896,971215059,807526948,778742000,586268941,365010934,951279875,316290802,267570670,583861472,851432142,435293607,286725742,722019349,8745084,730764433,739509517,470273943,209783453,680057396,889840849,569898238,459739080,29637311,489376391,519013702,8390086,527403788,535793874,63197655,598991529,662189184,261180706,923369890,184550589,107920472,292471061,400391533,692862594,93254120,786116714,879370834,665487541,544858368,210345902,755204270,965550172,720754435,686304600,407059028,93363621,500422649,593786270,94208912,687995182};
	return f[n]
}

```


![](http://wesub.ifree258.top/bottomPic.png)