package main

import "fmt"

// func mayPanic() {
//     panic("что-то пошло не так")
// }

// func main() {
//     fmt.Println("до panic")
//     mayPanic()
//     fmt.Println("после panic") // эта строка никогда не выполнится
// }


func mayPanic(){
	panic("деление на ноль невозможно")
}

func main(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановились после паники:",r)
		}
	}()




	var num1,num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num2 == 0{
		mayPanic()
	}

	result := num1 / num2
	fmt.Println("Результат деления:",result)
	fmt.Println("Программа продолжает работу")
}