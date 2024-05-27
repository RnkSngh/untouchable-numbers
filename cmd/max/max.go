package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var unsorted []int
	var maxNum = 0
	
	for scanner.Scan(){
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err!=nil {
			fmt.Println("Error converting string to integer:", err)
		return 
		}
		if num > maxNum{
			maxNum = num
		}
		unsorted = append(unsorted,  num)
	}

	fmt.Println("max number found", maxNum)
	sorted := make([]int, maxNum+1) 

	for _, integer := range unsorted {
		sorted[integer] = integer
	}
	fmt.Println("got sorted", sorted[0], sorted[1] , sorted[2])


	for i:=1; i< maxNum; i++{
		if sorted[i] != i{
			
			fmt.Println("max consecutive number found", i)
			return 
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
	
}


