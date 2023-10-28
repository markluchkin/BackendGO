package LinkedList

import "fmt"

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

func New(q int) *linkedList {
	// Инициализировать новый список. Количество элементов q.

	if q <= 0 {
		fmt.Println("ERROR: list size is incorrect.")
		return nil
	}

	head := node{}
	currNode := &head
	lst := linkedList{head: &head}

	for i := 0; i < q-1; i++ {
		currNode.next = &node{}
		currNode = currNode.next
	}
	return &lst
}

func (l *linkedList) Add(value int) {
	// Добавить новый элемент в конец.
	newNode := &node{val: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode

}

func (l *linkedList) Pop() {
	// Удалить элемент из конца.
	curr := l.head
	if curr.next == nil {
		return
	}

	for curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil
}

func (l *linkedList) At(pos int) int {
	// Получение элемента из позиции pos.
	if pos < 0 {
		fmt.Println("ERORR: inputted position doesn't exist.")
		return 0
	}
	curr := l.head
	i := 0
	for i != pos && curr.next != nil {
		curr = curr.next
		i++
	}

	return curr.val
}

func (l *linkedList) Size() int {
	// Вернуть длину списка.
	size := 0
	curr := l.head
	for curr.next != nil {
		size++
		curr = curr.next
	}
	size++
	return size
}

func (l *linkedList) DeleteFrom(pos int) {
	// Удалить элемент из позиции pos.
	curr := l.head
	if curr.next == nil {
		return
	}
	if pos == 0 {
		curr = curr.next
		return
	}

	for i := 0; i != pos-1 && curr.next != nil; i++ {
		curr = curr.next
	}

	curr.next = curr.next.next
}

func (l *linkedList) UpdateAt(pos, val int) {
	// Значение на позиции pos сделать равным val.
	if pos < 0 || pos >= l.Size() { // недопустимая позиция
		fmt.Println("ERORR: inputted position doesn't exist.")
		return
	}

	curr := l.head
	for i := 0; i < pos; i++ {
		curr = curr.next
	}

	curr.val = val
}
