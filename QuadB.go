package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {

	for i := 0; i < y; i++ { //rows
		if i == 0 {
			//for first row /***\
			for z := 0; z < x; z++ {
				// col
				if z == 0 {
					z01.PrintRune('/')
				} else if z == x-1 {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
		} else if i == y-1 {
			//for last row \***/
			for z := 0; z < x; z++ {
				// col
				if z == 0 {
					z01.PrintRune(92)
				} else if z == x-1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
		} else {
			// middle rows
			for z := 0; z < x; z++ {
				if z == 0 || z == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}

	}

}
