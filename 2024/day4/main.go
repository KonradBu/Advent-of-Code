package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	data, err := os.Open("/home/konradburgi/Documents/hhu/gitCode/adventOfCode/2024/day4/input")
	if err != nil{
		panic("file")
	}
	defer data.Close()
	scn := bufio.NewScanner(data)
	scn.Split(bufio.ScanLines)
	text := [][]rune{}
	lineCount := 0
	for scn.Scan(){
		line := []rune(scn.Text())
		text = append(text, []rune{})
		for _,r := range line{
			text[lineCount] = append(text[lineCount], r)
		} 
		lineCount++
	}
	sum := 0
	for i, l := range text{
			for j, c := range l{
				prevSum := sum 
				if c == 'X'{
					sum += check(text, 'M', 1,1,i,j)
					sum +=  check(text, 'M', 1,0,i,j)
					sum +=  check(text, 'M', 0,1,i,j)
					sum +=  check(text, 'M', 0,-1,i,j)
					sum +=  check(text, 'M', -1,0,i,j)
					sum +=  check(text, 'M', -1,-1,i,j)
					sum +=  check(text, 'M', -1,1,i,j)
					sum +=  check(text, 'M', 1,-1,i,j)
				}
				if prevSum < sum{
					//fmt.Println(i,j)
				}
			}
	}
	fmt.Println(sum)
}

func check(text [][]rune, next rune, dir1 int, dir2 int, x int, y int) int{
	checkx := x + dir1
	checky := y + dir2
	if checkx < 0|| checky < 0 || checkx >= len(text[y]) || checky >= len(text){
		return 0
	}
	if text[checkx][checky] == next{
		newNext := ' '
		switch next{
		case 'X':
			newNext = 'M'
		case 'M':
			newNext = 'A'
		case 'A':
			newNext = 'S'
		case 'S':
			fmt.Println(checkx, checky)
			return 1
		}
		return check(text, newNext, dir1, dir2, checkx, checky)
	}else{
		return 0
	}
}