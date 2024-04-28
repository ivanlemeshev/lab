package main

import (
	"fmt"
)

func main() {
	renderCube()
}

func renderCube() {
	size := 10

	leftOffset := 10
	rightOffset := 0

	for range size - 4 {
		for range leftOffset {
			fmt.Print(" ")
		}

		for range size * 2 {
			fmt.Print(".")
		}

		for range rightOffset {
			fmt.Print("0")
		}

		fmt.Print("\n")

		leftOffset -= 2
		rightOffset += 2
	}

	for range size - 6 {
		for range size * 2 {
			fmt.Print("1")
		}

		for range size {
			fmt.Print("0")
		}

		fmt.Print("\n")
	}

	rightOffset = 0

	for range size - 4 {
		for range size * 2 {
			fmt.Print("1")
		}

		for range size - rightOffset {
			fmt.Print("0")
		}

		fmt.Print("\n")

		rightOffset += 2
	}
}
