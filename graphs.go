package main

import "fmt"

// Queue Реализую тип очередь, так как в GO его нет.
type Queue struct {
	People  []Person
	Checked map[string]struct{}
}

// NewQueue возвращает экземпляр очереди
func NewQueue() *Queue {
	q := Queue{
		People:  make([]Person, 0, 100),
		Checked: map[string]struct{}{},
	}
	return &q
}

// Dequeue извлекает и возвращает первый элемент очереди.
func (q *Queue) Dequeue() (Person, error) {
	if len(q.People) < 1 {
		return Person{}, fmt.Errorf("queue is empty")
	}
	person := q.People[0]
	q.People = q.People[1:]
	return person, nil
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Enqueue(person ...Person) {
	for _, val := range person {
		q.People = append(q.People[:], val)
	}

}

// IsEmpty проверяет не пуста ли очередь.
func (q *Queue) IsEmpty() bool {
	return len(q.People) < 1
}

// Person тип - человек, который буду использовать для работы с графами
type Person struct {
	Name    string
	Friends []Person
}
