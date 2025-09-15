package main

import "fmt"

func AppendRange(min, max int) []int {

	if min >= max {
		return nil
	}
	result := max int; 
	result(min-1)

	return result
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))

}
