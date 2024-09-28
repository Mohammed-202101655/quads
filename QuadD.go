package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	for i := 0; i < y; i++ { //rows
		if i == 0 || i == y-1 {
			//first and last row (work on each col)
			for z := 0; z < x; z++ {
				if z == 0 {
					z01.PrintRune('A')
				} else if z == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else {
			//middle rows
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
