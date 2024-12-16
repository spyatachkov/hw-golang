package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}
type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Size int
	Head *ListItem
	Tail *ListItem
}

func NewList() List {
	return new(list)
}
func (l *list) Len() int {
	return l.Size
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	frontItem := l.Front()

	newItem := &ListItem{
		Value: v,
	}

	if frontItem == nil { // добавляется первый в список элемент
		// добавляемый элемент является и хвостом и головой
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Next = l.Head
		l.Head = newItem
		l.Head.Next.Prev = newItem
	}

	l.Size++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	backItem := l.Back()

	newItem := &ListItem{
		Value: v,
	}

	if backItem == nil {
		l.Tail = newItem
		l.Head = newItem
	} else {
		newItem.Prev = l.Tail
		l.Tail = newItem
		l.Tail.Prev.Next = newItem
	}
	l.Size++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	// удаление - это связывание предыдущего со следующем от удаляемого
	prev := i.Prev
	next := i.Next

	if prev == nil { // удаляется самый первый
		next.Prev = nil
		l.Head = next
	} else if next == nil { // удаляется самый последний
		prev.Next = nil
		l.Tail = prev
	} else {
		prev.Next = next
		next.Prev = prev
	}
	l.Size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}
