package main

import "fmt"

type epitech struct {
	ID          string
	Title       string 
	Description string
}

type allEpitech []epitech

var epitechs = allEpitech{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func main() {
	// nums := []int{2,3,4, 12 ,4, 56 ,6, 8, 9, 675}



	
	fmt.Println(epitechs[0:])
	// fmt.Println(nums[1:])
}