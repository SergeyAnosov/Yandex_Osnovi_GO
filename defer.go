package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("Hello")
	//for i := 1; i <= 3; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("World")

	//fmt.Println(unintuitive())
	//fmt.Println(intuitive())

	//VeryLongTimeFunction()
	fmt.Println(Global)
	UseGlobal()
	fmt.Println(Global)

}

var Global = 5

func UseGlobal() {
	defer func(checkout int) {
		Global = checkout // присваиваем Global значение аргумента
	}(Global) // копируем значение Global в аргументы функции
	Global = 42 // Изменяем Global
	fmt.Println(Global)
	// Здесь будет вызвана отложенная функция
}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
	// Какие-то долгие вычисления
}
