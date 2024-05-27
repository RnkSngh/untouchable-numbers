package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func printNumber(a int){
	fmt.Println(strings.Repeat(strconv.Itoa(a), a))
	fmt.Println("printed", a)
}


func wait(){
	time.Sleep(3 * time.Second)
}

func main(){
	c := time.After(3*time.Second) 

	go func() {
		for i:=0; true; i++{
			go printNumber(i)
		}
	}()
	
	<-c
}
