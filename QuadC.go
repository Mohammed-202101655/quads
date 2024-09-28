package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	for i := 0; i < y; i++ { //rows
		if i == 0 {
			//first row (work on each col)
			for z := 0; z < x; z++ {
				if z == 0 || z == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else if i == y-1 {
			//last row (work on each col)
			for z := 0; z < x; z++ {
				if z == 0 || z == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else {
			for z := 0; z < x; z++ {
				if z == 0 || z == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
