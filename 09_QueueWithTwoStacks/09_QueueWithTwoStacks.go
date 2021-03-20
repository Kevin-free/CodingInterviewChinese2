package _9_QueueWithTwoStacks

type CQueue struct {
	in  stack
	out stack
}

type stack []int

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
