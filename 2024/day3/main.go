package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	data, err := os.Open("/home/konradburgi/Documents/hhu/gitCode/adventOfCode/2024/day3/input.txt")
	if err != nil{
		panic("File :(")
	}
	defer data.Close()
	scn := bufio.NewScanner(data)
	scn.Split(bufio.ScanRunes)
	sum := 0
	buffer := ""
	var val1 int
	var val2 int
	var s1 string
	var s2 string
	do := true
	for scn.Scan(){
		r := scn.Text()
		_, err := strconv.Atoi(r)
		var last string
		if len(buffer) == 0{
			last = ""
		}else{
			last = string([]rune(buffer)[len(buffer) -1])
		}
		
		switch {
		case r == "m":
			buffer = "m"
		case r == "u" && buffer == "m":
			buffer = "mu"
		case r == "l" && buffer == "mu":
			buffer = "mul"
		case r == "d":
			buffer = "d"
		case r == "o" && buffer == "d":
			buffer = "do"
		case r == "("  && buffer == "mul":
				buffer = "mul(#"
		case r == "(" && buffer == "do":
				buffer = "do("
		case r == "(" && buffer == "don,t":
				buffer = "don,t("
		case r == "n" && buffer == "do":
			buffer = "don"
		case r == "'" && buffer == "don":
			buffer = "don,"
		case r == "t" && buffer == "don,":
			buffer = "don,t"
		case last == "#" && err == nil:
			s1 += r
			buffer += r + "#"
		case r == "," && last == "#":
			buffer += "!" 
		case last == "!" && err == nil:
			s2 += r
			buffer += r + "!"
		case r == ")" && last == "!":
			val1, _ = strconv.Atoi(s1)
			val2, _ = strconv.Atoi(s2)
			if do{
				sum += val1 * val2
			}
			buffer = ""
			s1 = ""
			s2 = ""
			val1 = 0
			val2 = 0
		case r == ")" && last == "(":
			fmt.Println(buffer)
			if buffer == "don,t("{
				do = false
			}
			if buffer == "do("{
				do = true
			}
			buffer = ""
			s1 = ""
			s2 = ""
			val1 = 0
			val2 = 0
		default:
			buffer = ""
			buffer = ""
			s1 = ""
			s2 = ""
			val1 = 0
			val2 = 0
		}
	}
	fmt.Println(sum)
}