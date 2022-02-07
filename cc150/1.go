package main

import "fmt"

func main() {
	str := "abcfdefa"
	fmt.Println(isUnique(str))
}

func isUnique(str string) bool {
	bitmap := 0
	for _, r := range str {
		pos := r - 'a'
		fmt.Printf("[%q] bitmap(%d) & %d (1 << %d) = %d\t", r, bitmap, 1<<pos, pos, bitmap&(1<<pos))
		if (bitmap & (1 << pos)) != 0 {
			return false
		}
		bitmap |= (1 << pos)
		fmt.Printf("current-bitmap = %d\n", bitmap)
	}
	return true
}
