package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type linkedList struct {
	Head *Node
}

func New(q int) *linkedList {
	// Инициализировать новый список. Количество элементов q.
	head := Node{}
	currNode := &head
	lst := linkedList{Head: &head}
	for i := 0; i < q; i++ {
		currNode.Next = &Node{}
		currNode = currNode.Next
	}
	return &lst
}

func (l *linkedList) Add(val int) {
	// Добавить элемент в конец списка.
	newNode := &Node{Val: val, Next: nil}

	if l.Head == nil {
		l.Head = newNode
	} else {
		currNode := l.Head
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = newNode
	}
}

func (l *linkedList) Pop() {
	// Удалить элемент из конца.
	curr := l.Head
	for i := 0; i < l.Size()-2; i++ {
		curr = curr.Next
	}
	curr.Next = nil
}

func (l *linkedList) At(pos int) int {
	// Получение элемента из позиции pos.
	curr := l.Head
	i := 0
	for i != pos {
		curr = curr.Next
		i++
	}

	return curr.Val
}

func (l *linkedList) Size() int {
	// Вернуть длину списка.
	size := 0
	curr := l.Head
	for curr.Next != nil {
		size++
		curr = curr.Next
	}

	return size
}

func (l *linkedList) DeleteFrom(pos int) {
	// Удалить элемент из позиции pos.
	if pos < 0 || pos >= l.Size() { // недопустимая позиция
		return
	}

	if pos == 0 { // то удаляем 1 элемент списка
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	prev := curr
	for i := 0; i < pos; i++ {
		prev = curr
		curr = curr.Next
	}

	prev.Next = curr.Next
}

func (l *linkedList) UpdateAt(pos, val int) {
	// Значение на позиции pos сделать равным val.
	if pos < 0 || pos >= l.Size() { // недопустимая позиция
		return
	}

	curr := l.Head
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}

	curr.Val = val
}

//NewFromSlice *

func main() {
	list := New(10)
	list.Add(5)         // 11 эл
	list.Add(131231231) // 12 эл
	fmt.Println(list.Size())
	fmt.Println(list.At(12))
	fmt.Println(list.At(0))
	list.DeleteFrom(11)
	fmt.Println(list.At(11))

}
