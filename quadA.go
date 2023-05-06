package main

import "github.com/01-edu/z01"

func main() {
<<<<<<< HEAD
	quadA(0, 1)
=======
	quadA(1, 0)
>>>>>>> 809e685 (ajout des fichiers)
}

func quadA(x, y int) {
	for h := 1; h <= y; h++ {
		for l := 1; l <= x; l++ {
			if h == 1 || h == y {
				if l == 1 || l == x {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if l == 1 || l == x {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
