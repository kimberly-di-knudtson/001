package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	size := 10000
	k := rand.Intn(3 * size)
	num_list := make([]int, size)
	for i := range num_list {
		num_list[i] = rand.Intn(size)
	}

	start := time.Now()
	v1, v2, sum := twoSumK_nsquared(k, num_list)
	fmt.Printf("~n-squared took %v us\n", time.Since(start).Microseconds())
	fmt.Printf("%v, %v, %v+%v, %v\n", k, len(num_list), v1, v2, sum)

	start = time.Now()
	v1, v2, sum = twoSumK_n(k, num_list)
	fmt.Printf("~n took %v us\n", time.Since(start).Microseconds())
	fmt.Printf("%v, %v, %v+%v, %v\n", k, len(num_list), v1, v2, sum)
}

func twoSumK_nsquared(k int, num_list []int) (int, int, bool) {
	//n-squared - check each sum combination
	ops := 0
	for i1, v1 := range num_list {
		for i2, v2 := range num_list[i1+1:] {
			sum := 0
			if i1 != i2 {
				sum = v1 + v2
				ops++
				if sum == k {
					fmt.Printf("num ops = %v\n", ops)
					return v1, v2, true
				}
			}
		}
	}
	fmt.Printf("num ops = %v\n", ops)
	return 0, 0, false
}

func twoSumK_n(k int, num_list []int) (int, int, bool) {
	//n - sort first, then work from the ends
	ops := 0
	start := time.Now()
	sort.Ints(num_list)
	fmt.Printf("sorting took: %v us\n", time.Since(start).Microseconds())

	min_index := 0
	max_index := len(num_list) - 1
	for {
		sum := num_list[max_index] + num_list[min_index]
		ops++
		if max_index == min_index {
			break
		}
		if sum > k {
			max_index--
		} else if sum < k {
			min_index++
		} else {
			fmt.Printf("num ops = %v\n", ops)
			return num_list[min_index], num_list[max_index], true
		}

	}
	fmt.Printf("num ops = %v\n", ops)
	return 0, 0, false
}
