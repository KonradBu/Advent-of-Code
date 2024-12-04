package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("/home/konradburgi/Documents/hhu/gitCode/adventOfCode/2024/two/errorupdate.txt")
	if err != nil {
		panic(" File no open :(")
	}
	defer data.Close()

	sum1 := 0
	sum2 := 0
	scn := bufio.NewScanner(data)
	for scn.Scan() {
		safe1 := safe(conv(scn.Text()), true)
		safe2 := safe2(conv(scn.Text()), true)
		sum1 += safe1
		sum2 += safe2
		if safe1 == safe2 {
		} else {
			fmt.Println(scn.Text())
			fmt.Println(safe1)
			fmt.Println(safe2)
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func conv(input string) []int {
	split := strings.Fields(input)
	arr := []int{}
	for _, s := range split {
		val, _ := strconv.Atoi(s)
		arr = append(arr, val)
	}
	return arr
}

func safe2(arr []int, tolerate bool) int {
	increasing := 1
	var prev int
	for i, num := range arr {
		if i == 0 {
			prev = num
			continue
		}
		if i == 1 {
			if num < prev {
				increasing = -1
			}
		}

		res := (num - prev) * increasing

		if res <= 0 {
			if tolerate {
				if i == 2 {
					if safe2(arr[1:], false) == 1 {
						return 1
					}
					/*a := append(arr[:1], arr[2:]...)
					if safe2(a, false) == 1{
						return 1
					}
					*/
				}
				arr2 := make([]int, len(arr))
				copy(arr2, arr)
				if (safe2(append(arr2[:i -1], arr2[i:]...), false)) == 1{
					return 1
				}
				arr3 := make([]int, len(arr))
				copy(arr3, arr)
				a := safe2(append(arr3[:i], arr3[i+1:]...), false)
				return a
			} else {
				return 0
			}
		}
		if res > 3 {
			if tolerate {
				if safe2(arr[1:], false) == 1 {
					return 1
				}
				return safe2(append(arr[:i], arr[i+1:]...), false)
			} else {
				return 0
			}
		}
		prev = num
	}
	return 1
}

func safe(arr []int, t bool) int {
	increasing := 1
	prev := arr[0]
	tolerate := t
	for i, num := range arr {
		if i == 0 {
			prev = num
			continue
		}
		if i == 1 {
			if prev > num {
				increasing = -1
			}
		}
		res := (num - prev) * increasing
		if res <= 0 {
			if tolerate {
				if i == 2 {
					if safe(arr[1:], false) == 1 {
						return 1
					}
					if safe(append(arr[:1], arr[2:]...), false) == 1 {
						return 1
					}
				}
				return safe(append(arr[:i], arr[i+1:]...), false)
			} else {
				return 0
			}
		} else if res > 3 {
			if tolerate {
				tolerate = false
			} else {
				return 0
			}
		} else {
			prev = num
		}
	}
	return 1
}
