package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeDuplicates(nums []int) int {
	var k int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] && i > k {
			if i > k+1 {
				nums[k+1] = nums[i]
			}
			k++
		}
	}
	return k + 1
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	lines := strings.Fields(line)

	var num int
	nums := make([]int, len(lines))
	for i, s := range lines {
		num, _ = strconv.Atoi(s)
		nums[i] = num
	}

	fmt.Print(removeDuplicates(nums))

}
