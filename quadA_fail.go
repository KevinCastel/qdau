package main

import "github.com/01-edu/z01"

func main() {
	QuadA(1, 1)
}

func QuadA(x, y int) {
	// y est le nombre de cote
	// x est le nombre de charactère

	for c := x; c <= x+2; c++ {
		for l := y; l <= y+4; l++ {
			if c == x+1 || c == x+4 {
				if l == y || l == y+4 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
				/*
					if l > y && l < y+5 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('|')
					} */
			} else if l == y || l == y+4 {
				if c == x || c == x+2 {
					z01.PrintRune('o')
				}
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
	}
}

/*
func QuadA(x, y int) {
	for y := 1; y <= 4; y++ {
		for x := 1; x <= 5; x++ {
			if y == 1 {
				if x == 1 || x == 5 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else if y == 2 {
				if x ==
			}


			if y == 2 || y == 4 {
				z01.PrintRune('|')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('-')
			}
		}
	}

	if x > 0 && y > 00 {
	}
}
*/
