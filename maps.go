package main

import "fmt"

func main() {
	porducts := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}
	var list = make([]string, 0, 10)
	var sum int = 0
	for k, v := range porducts {

		if v > 500 {
			sum += v
			list = append(list, k)
		}
	}
	fmt.Println(list)
	fmt.Println("sum = ", sum)

}
