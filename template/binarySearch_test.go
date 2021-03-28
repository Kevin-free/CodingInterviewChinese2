package template

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 9}
	target := 4
	res := binarySearchWithR(nums, target)
	fmt.Printf("res: %v\n", res)
}
