package main

import (
	"fmt"
)

const num = 8

func main() {
	var ans [num]int
	var sum = 0
	var i, j int
	for i = 0; i < num; i++ {
		ans[i] = -1
	}
	print("begin\n")
	for i = 0; i > -1; {
		for j = ans[i] + 1; true; {
			if j >= num {
				i--
				break
			}
			if check(ans, i, j) {
				ans[i] = j
				if i == num-1 {
					prtans(ans)
					sum++
				} else {
					i++
					ans[i] = -1
					break
				}
			}
			j++
		}
	}
	fmt.Printf("sum=%d", sum)
}

func check(ans [num]int, x, y int) bool {
	var i int
	for i = 0; i < x; i++ {
		if ans[i] == y || ans[i]-i == y-x || ans[i]+i == y+x {
			return false
		}
	}
	return true
}

func prtans(ans [num]int) {
	var i, j int
	for i = 0; i < num; i++ {
		for j = 0; j < num; j++ {
			if j == ans[i] {
				print("*")
			} else {
				print("0")
			}
		}
		print("\n")
	}
	print("--------------\n")
}
