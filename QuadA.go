package piscine

import "github.com/01-edu/z01"

func QuadA(x,y int) {
	for i := 0 ; i < y ; i++ { //row
		// first or last row o---o
		if i == 0 || i == y-1{
			for z := 0 ; z < x; z++ {
				if z == 0 || z == x-1{
					z01.PrintRune('o')
				}else{
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		}else{
			for z := 0 ; z < x; z++ { //column
				//first and last | in middle rows
				if z == 0 || z == x-1{
					z01.PrintRune('|')
				}else{
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')

		}
	}
		
}
	