package main

import "fmt"

func fibonacci() func() int {
	f := []int{0, 1}
	count := 0
	return func() (value int) {
		if count == len(f) {
			f = append(f, f[count-1]+f[count-2])
		}
		value = f[count]
		count += 1
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
