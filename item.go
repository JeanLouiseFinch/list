package list

import (
	"errors"
)

//Item элемент списка
type Item struct {
	value interface{}
	next *Item
	prev *Item
}

//NewItem возвращает указатель на новый элемент
func NewItem(value interface{}) *Item{
	return &Item{value:value}
}

//Next возвращает указатель на следующий элемент
func (i *Item) Next() *Item {
		return i.next
}

//Prev возвращает указатель на предыдущий элемент
func (i *Item) Prev() *Item {
	return i.prev
}

//Value получение значения. 
func (i *Item) Value() (interface{},error) {
	if i == nil {
		return nil,errors.New("Element is nil")
	}
	return i.value,nil
}
