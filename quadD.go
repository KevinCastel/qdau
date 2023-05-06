package main

import "github.com/01-edu/z01"

func main() {
	QuadD(8, 10)
}

func QuadD(x, y int) {
	for v := 1; v <= y; v++ {
		for h := 1; h <= x; h++ {
			if v == 1 {
				if h == 1 {
					z01.PrintRune('A')
				} else if h == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if v == y {
				if h == 1 {
					z01.PrintRune('A')
				} else if h == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if h == 1 || h == x {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
