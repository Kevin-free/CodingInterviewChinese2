# [剑指 Offer 09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)


## 题目

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )


示例 1：
```
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
```

示例 2：
```
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
```
提示：
```
1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
```


## 解题思路

这道题要求我们使用两个"先进后出"的栈实现一个"先进先出"的队列。

我们可以将一个栈负责插入操作，另一个栈负责删除操作。

根据栈先进后出的特性，我们每次往第一个栈里插入元素后，第一个栈的顶部元素是最后插入的元素，第一个栈的顶部元素是下一个待删除的元素。
为了维护队列先进先出的特性，第二个栈就派上用场了。
用第二个栈维护待删除的元素，在执行删除操作的时候我们首先看下第二个栈是否为空。
如果为空，我们将第一个栈里的元素一个个弹出插入到第二个栈里，这样第二个栈里元素的顺序就是待删除的元素的顺序，
要执行删除操作的时候我们直接弹出第二个栈的元素返回即可。


![0901.gif](https://pic.leetcode-cn.com/1616241111-yHSPfH-0901.gif)


## 代码

```go
type CQueue struct {
	in  stack
	out stack
}

type stack []int // define stack

func (s *stack) Push(value int) {
	*s = append(*s, value)
}

func (s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1] // get the last one
	*s = (*s)[:n-1]  // pop the last one
	return res
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) != 0 {
		// out 中有，直接 pop
		return this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			// out 中没有，但 in 中有，将in中所有元素移至out中
			this.out.Push(this.in.Pop())
		}
		// 在out中pop出一个
		return this.out.Pop()
	}
	// out & in 中都没有，return -1
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
```


![](http://wesub.ifree258.top/bottomPic.png)