package main

import (
	"BackendGo/LinkedList"
	"fmt"
)

func main() {
	// проверка на корректную работоспособность всех функций
	list := LinkedList.New(5)
	fmt.Println(list.Size()) // output 5
	list.Add(1234)
	list.Add(12345)
	fmt.Println(list.Size()) // output 7
	list.Pop()
	fmt.Println(list.Size())            // output 6
	fmt.Println(list.At(1), list.At(6)) // output 0 1234
	list.Add(123456)
	//list.Add(12323123)
	list.DeleteFrom(6)
	fmt.Println(list.At(6)) // output 123456
	list.UpdateAt(0, 1111)
	list.UpdateAt(1, 2222)
	fmt.Println("values of list:")
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("pos = %d, val = %d \n", i, list.At(i))
	}
	fmt.Println(list.Size()) //output 7
}
