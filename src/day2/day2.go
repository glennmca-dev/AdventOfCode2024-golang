package main

import (
	"AdventOfCode2024/util"
	"fmt"
)

func main(){
	lines, err := util.ReadFileLines("./day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
