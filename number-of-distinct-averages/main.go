package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	sort.Ints(nums)

	averages := make(map[int]int)
	for i := 0; i < len(nums)/2; i++ {
		ave := int((float64(nums[i]+nums[len(nums)-i-1]) / 2.0) * 10)
		averages[ave]++
	}

	fmt.Print(len(averages))

}
