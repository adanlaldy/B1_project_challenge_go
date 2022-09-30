package main

import (
	piscine "exam1/atoi"
	"os"
)

func b(parlenght, parlength int) {
	corner := "X"
	lenght := "-"
	spaces := " "
	incre := 0
	decre := -2
	if parlenght == 1 && parlength == 1 {
		println(corner)
	} else if parlenght == 1 && parlength == 3 {
		print(corner)
		print(lenght)
		println(corner)
	} else if parlenght == 3 && parlength == 1 {
		println(corner)
		println(0)
		println(corner)
	} else {
		for i := 0; i < 1; i++ {
			print(corner)
			for j := 0; j < parlenght-2; j++ {
				print(lenght)
			}
			println(corner)
		}
		for k := 2; k < parlength; k++ {
			print(parlength - parlength + incre)
			for l := 1; l < parlenght-1; l++ {
				print(spaces)
			}
			a := parlength + decre
			if a < 10 {
				println(a)
			} else if a == 10 {
				println(0)
			} else if a > 10 {
				println(a - 10)
			}
			decre++
			incre++
		}
		for m := 0; m < 1; m++ {
			print(corner)
			for m := 0; m < parlenght-2; m++ {
				print(lenght)
			}
			println(corner)
		}
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 || len(arg) == 1 {
		print("incomplete...")
		return
	} else if piscine.Atoi(arg[0]) == 0 || piscine.Atoi(arg[1]) == 0 {
		print("not possible...")
		return
	} else {
		b(piscine.Atoi(arg[0]), piscine.Atoi(arg[1]))
	}
}
