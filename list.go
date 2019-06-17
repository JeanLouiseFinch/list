package list

import (
	"sync"
	"errors"
	"fmt"
)

//List структура "список". для обновременного доступа добавлен мьютекс
type List struct{
	first *Item
	last *Item
	len int
	sync.Mutex
}

//NewList Создает новый список
func NewList() *List {
	return &List{}
}


//Print для удобного просмотра результата
func (l *List) Print() {

	if l.len == 0 {
		fmt.Println("List empty!")
		return
	}
	
	for it := l.first; it != nil; it = it.next {
		if it.prev != nil {
			fmt.Printf(" <-> (%v)",it.value)
		} else {
			fmt.Printf("\tList: (%v)",it.value)
		}
	}	
	fmt.Printf("\n")
}
//First возвращает первый элемент или ошибку если его нет
func (l *List) First() (*Item,error) {
	if l.first == nil {
		return nil,errors.New("First element nil")
	}
	return l.first,nil
}

//Last возвращает последний элемент или ошибку если его нет
func (l *List) Last() (*Item,error) {
	if l.last == nil {
		return nil,errors.New("Last element nil")
	}
	return l.last,nil
}

//Len возвращает длину списка
func (l *List) Len() int{
	return l.len
}

//PushFront вставка нового значения в начале списка
func (l *List) PushFront(value interface{}) {
	l.Lock()
	new := NewItem(value)
	if l.first != nil {
		new.next = l.first
		l.first.prev = new
	}
	if l.last == nil {
		l.last = new
	}
	l.first = new
	l.len++
	l.Unlock()
}

//PushBack вставка нового значения в конце списка
func (l *List) PushBack(value interface{}) {
	l.Lock()
	new := NewItem(value)
	if l.last != nil {
		new.prev = l.last
		l.last.next = new
	}
	if l.first == nil {
		l.first = new
	}
	l.last = new
	l.len++
	l.Unlock()
}

// RemoveItem удаляем элемент. очищаем все ссылки
func (l *List) RemoveItem(ri *Item) {
	l.Lock()
	if ri == nil || l.len == 0 {
		return
	}
	
	if ri.prev != nil {
		ri.prev.next = ri.next
	} else {
		if l.first == ri {
			l.first = ri.next
		}
	}
	if ri.next != nil {
		ri.next.prev = ri.prev
	} else {
		if l.last == ri {
			l.last = ri.prev
		}
	}

	ri.next = nil
	ri.prev = nil

	l.len--
	l.Unlock()
}