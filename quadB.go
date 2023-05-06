package main

import "github.com/01-edu/z01"

func main() {
	QuadB(5, 5)
}

func QuadB(x, y int) {
	for v := 1; v <= y; v++ {
		for h := 1; h <= x; h++ {
			if v == 1 {
				if h == 1 {
					z01.PrintRune('/')
				} else if h == x {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else if v == y {
				if h == 1 {
					z01.PrintRune('\\')
				} else if h == x {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else if h == 1 || h == x {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
