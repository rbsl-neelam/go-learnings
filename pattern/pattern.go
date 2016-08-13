package pattern

import (
	"fmt"
)


func Triangle1 (n int) {
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			if j == 0 || i == n - 1 || i == j{
				fmt.Print("*")
			} else{
				fmt.Print(" ")
			}
		}
		fmt.Println("\n")
	}
}


func Triangle2 (n int) {
	for i := 0; i < n; i++{
		for j := n - 1; j >= 0; j--{
			if j == 0 || i == n - 1 || i == j{
				fmt.Print("*")
			} else{
				fmt.Print(" ")
			}
		}
		fmt.Println("\n")
	}
}


func Triangle3 (n int) {
	for i := n - 1; i >= 0; i--{
		for j := 0; j < n; j++{
			if i == n - 1 || j == 0 || i == j{
				fmt.Print("*")
			} else{
				fmt.Print(" ")
			}
		}
		fmt.Println("\n")
	}
}


func Triangle4 (n int) {
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			if i == 0 || j == n - 1 || i == j{
				fmt.Print("*")
			} else{
				fmt.Print(" ")
			}
		}
		fmt.Println("\n")
	}
}

