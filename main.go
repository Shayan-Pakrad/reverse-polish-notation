package main

import (
	"fmt"
	"strconv"
)

func main() {

	statement := []string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"} // given statement
	var stack []int                                                                                   // results stack

	for len(statement) > 0 {

		if statement[0] == "+" {
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, result)
			statement = statement[1:]
			continue
		}

		if statement[0] == "-" {
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, result)
			statement = statement[1:]
			continue
		}

		if statement[0] == "/" {
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, result)
			statement = statement[1:]
			continue
		}
		if statement[0] == "*" {
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, result)
			statement = statement[1:]
			continue
		}

		number, err := strconv.Atoi(statement[0])
		if err != nil {
			fmt.Println(err)
		}
		stack = append(stack, number)
		statement = statement[1:]
	}

	fmt.Println(stack[0])
}
