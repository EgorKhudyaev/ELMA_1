package main

import (
	"fmt"
)

func main() {
	// Переменные для функции
	var z = 0
	for s := 1; s < 101 ; s++ {
		z = 0
		if s%3 == 0 {z = 1} 
		if s%5 == 0 {z = 2}
		if s%15 == 0 {z = 3}
		switch z {
		case 1:
        		fmt.Println("Fizz")
		case 2:
        		fmt.Println("Buzz")
		case 3:
        		fmt.Println("FizzBuzz")
		case 0:
        		fmt.Println(s)
    		}
	}
}
