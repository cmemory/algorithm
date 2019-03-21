package datastructure

import "fmt"

type Queue struct {
	head *LNode
	tail *LNode
	len int
}

func NewQueue() *Queue{
	return &Queue{nil,nil,0}
}

func (q *Queue) QLen() int{
	return q.len
}

func (q *Queue) IsEmpty() bool{
	return q.len == 0
}

func (q *Queue) Push(val interface{}){
	node := &LNode{val,nil}
	if q.len == 0 {
		q.head = node
	}else {
		q.tail.Next = node
	}
	q.tail = node
	q.len ++
}

func (q *Queue) pop() (interface{},bool){
	if q.len == 0 {
		fmt.Println("queue is empty!")
		return nil,false
	}
	tmp := q.head
	q.head = tmp.Next
	tmp.Next = nil
	q.len --
	if q.head == nil {
		q.tail = nil
	}
	return tmp.T,true
}
