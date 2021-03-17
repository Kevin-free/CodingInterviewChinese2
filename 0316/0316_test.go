package _316

import (
	"fmt"
	"testing"
)

func Test0316(t *testing.T) {
	input := []int{1,3,4,5,8,8,8,11,18}
	output := searchLeft(input, 3)
	fmt.Println("output:",output)
	output = searchLeft(input, 8)
	fmt.Println("output:",output)
	output = searchLeft(input, 10)
	fmt.Println("output:",output)
}
