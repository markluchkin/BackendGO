package main

import (
	"BackendGo/LinkedList"
	"fmt"
)

func main() {
	// проверка на корректную работоспособность всех функций
	list := LinkedList.New(5)
	fmt.Println(list.Size())
	list.UpdateAt(0, 1)
	list.UpdateAt(1, 2)
	list.UpdateAt(2, 3)
	list.UpdateAt(3, 4)
	list.UpdateAt(4, 5)
	fmt.Println(list.Size())
	list.Add(6)
	fmt.Println(list.Size())
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("pos = %d, val = %d \n", i, list.At(i))
	}
	list.Pop()
	fmt.Println(list.Size())
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("pos = %d, val = %d \n", i, list.At(i))
	}
	//list.DeleteFrom(0)
	list.DeleteFrom(1)
	fmt.Println(list.Size())
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("pos = %d, val = %d \n", i, list.At(i))
	}
	list.DeleteFrom(1)
	fmt.Println(list.Size())
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("pos = %d, val = %d \n", i, list.At(i))
	}

}
