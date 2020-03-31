package main

import "fmt"

func main(){
	// for
	for num:= 0; num < 10; num++{
		fmt.Println("FOR: ")
		fmt.Println(num)
	}

	// while
	for num:= 0; num < 10;{
		fmt.Println("WHILE: ")
		fmt.Println(num)
		num++
	}
}