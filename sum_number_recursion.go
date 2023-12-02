package main

import "fmt"

func main() {
	fmt.Println(sumn(5))
}

func sumn(n uint) uint {
	if n < 0 { // терминальная ветка
		return 0
	} else { // рекурсивная ветка
		return n + sumn(n-1)
	}
}
