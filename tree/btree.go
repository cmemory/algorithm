package tree

import (
	"algorithm/datastructure"
	"fmt"
)

type BTree struct {
	root *datastructure.BNode
}

// 以先序遍历的字符串构建二叉树
// 如：AB##CD##EF##G#H##
func Create(str string, i int, len int) (*datastructure.BNode, int) {
	if i > len-1 || str[i] == '#'{
		return nil,i+1
	}
	left,index := Create(str,i+1,len)
	right,index := Create(str,index,len)
	return &datastructure.BNode{T:str[i],Left:left,Right:right},index
}

func CreateBTree(str string) *BTree{
	node,_ := Create(str,0,len(str))
	return &BTree{node}
}

func (bt *BTree) PreOrder() {
	if bt.root == nil {
		return
	}
	fmt.Printf("%c", bt.root.T)
	(&BTree{bt.root.Left}).PreOrder()
	(&BTree{bt.root.Right}).PreOrder()
}

func (bt *BTree) PreOrder2() {
	if bt.root == nil {
		return
	}
	p := bt.root
	s := datastructure.NewStack()
	for p!=nil || !s.IsEmpty() {
		if p!=nil {
			s.Push(p)
			fmt.Printf("%c",p.T)
			p = p.Left
		}else{
			t,_ := s.Pop()
			p = t.(*datastructure.BNode).Right
		}
	}
	fmt.Println()
}

func (bt *BTree) PreOrder3() {
	if bt.root == nil {
		return
	}
	s := datastructure.NewStack()
	s.Push(bt.root)
	p := bt.root
	// root节点类似于哨兵，弹出root时刚好栈空退出循环
	for !s.IsEmpty() {
		fmt.Printf("%c",p.T)
		if p.Right != nil {
			s.Push(p.Right)
		}
		if p.Left != nil{
			p = p.Left
		}else {
			t, _ := s.Pop()
			p = t.(*datastructure.BNode)
		}
	}
	fmt.Println()
}

//func (bt *BTree) PreOrder3() {
//	if bt.root == nil {
//		return
//	}
//	s := datastructure.NewStack()
//	s.Push(bt.root)
//	p := bt.root
//	for !s.IsEmpty() {
//		t, _ := s.Pop()
//		p = t.(*datastructure.BNode)
//		fmt.Printf("%c",p.T)
//		if p.Right != nil {
//			s.Push(p.Right)
//		}
//		if p.Left != nil{
//			s.Push(p.Left)
//		}
//	}
//	fmt.Println()
//}

func (bt *BTree) MorrisPreOrder(){
	if bt.root==nil {
		return
	}
	r := bt.root
	p := bt.root
	for r!=nil {
		p = r.Left
		if p != nil {
			for p.Right!=nil && p.Right!=r{
				p = p.Right
			}
			if p.Right == nil {
				fmt.Printf("%c",r.T)
				p.Right = r
				r = r.Left
				continue
			}else{
				p.Right = nil
			}
		}else {
			fmt.Printf("%c",r.T)
		}
		r = r.Right
	}
	fmt.Println()
}

func (bt *BTree) MorrisInOrder(){
	if bt.root == nil {
		return
	}
	r := bt.root
	p := bt.root
	for r != nil {
		p = r.Left
		if p != nil {
			for p.Right != nil && p.Right != r {
				p = p.Right
			}
			if p.Right == nil {
				p.Right = r
				r = r.Left
				continue
			}else {
				p.Right = nil
			}
		}
		fmt.Printf("%c",r.T)
		r = r.Right
	}
	fmt.Println()
}

func (bt *BTree) InOrder()  {
	if bt.root == nil {
		return
	}
	(&BTree{bt.root.Left}).InOrder()
	fmt.Printf("%c",bt.root.T)
	(&BTree{bt.root.Right}).InOrder()
}

func (bt *BTree) InOrder2()  {
	if bt.root == nil {
		return
	}
	p := bt.root
	s := datastructure.NewStack()
	for p!=nil || !s.IsEmpty() {
		if p!=nil {
			s.Push(p)
			p = p.Left
		}else{
			t,_ := s.Pop()
			p = t.(*datastructure.BNode)
			fmt.Printf("%c",p.T)
			p = p.Right
		}
	}
	fmt.Println()
}

func (bt *BTree) PostOrder()  {
	if bt.root == nil {
		return
	}
	(&BTree{bt.root.Left}).PostOrder()
	(&BTree{bt.root.Right}).PostOrder()
	fmt.Printf("%c",bt.root.T)
}

func (bt *BTree) PostOrder2(){
	if bt.root == nil {
		return
	}
	p := bt.root
	tmpS := datastructure.NewStack()
	outS := datastructure.NewStack()
	for p != nil || !tmpS.IsEmpty() {
		if p != nil {
			tmpS.Push(p)
			outS.Push(p)
			p = p.Right
		}else{
			t,_ := tmpS.Pop()
			p = t.(*datastructure.BNode).Left
		}
	}
	for !outS.IsEmpty() {
		p,_ := outS.Pop()
		fmt.Printf("%c",p.(*datastructure.BNode).T)
	}
	fmt.Println()
}

func (bt *BTree) PostOrder3(){
	if bt.root == nil {
		return
	}
	cur := bt.root
	last := bt.root
	s := datastructure.NewStack()
	for cur!=nil || !s.IsEmpty() {
		for cur!=nil  {
			s.Push(cur)
			cur = cur.Left
		}
		t,_ := s.Top()
		cur = t.(*datastructure.BNode)
		// (cur.Left==nil&&cur.Right==nil) || (cur.Right!=nil&&last==cur.Right) || (cur.Right==nil&&cur.Left!=nil&&last==cur.Left)
		// 1、当循环开始cur!=nil时，由前面for cur!=nil循环可知，最终栈顶cur的Left必为nil。所以第一个条件等价于cur.Right==nil ，第三个条件总为false
		// --> cur.Right==nil || (cur.Right!=nil&&last==cur.Right)
		// --> 看成 A || (~A && B),利用集合公式化简为 (A || ~A)&&(A||B) --> A||B
		// 也可这样分析：当前一个条件为true时 整个为true，当前一个为false时，后一个等价于last==cur.Right
		// --> cur.Right==nil || last==cur.Right
		// 2、当循环开始cur==nil时，需要取栈顶元素作为cur，由之前分析可知左子树访问完了且右边未访问。
		// 显然(cur.Left==nil||(cur.Left!=nil&&last==cur.Left))为true
		// 若cur.Left==nil，分析同1.
		// 若cur.Left!=nil&&last==cur.Left，第一个条件为false，第二个条件也为false，第三个条件等价于cur.Right==nil
		//
		// 综上：最终条件为 cur.Right==nil || last==cur.Right
		// 前面条件表达式实际上有几处是可以利用&&和||的短路特性进行简化的，简化之后再来讨论相应会简单很多。
		if cur.Right==nil || last==cur.Right{
			s.Pop()
			fmt.Printf("%c",cur.T)
			last = cur
			cur = nil
		}else{
			cur = cur.Right
		}
	}
	fmt.Println()
}

func (bt *BTree) Post(){
	if bt.root == nil {
		return
	}
	s := datastructure.NewStack()
	cur := bt.root
	var last *datastructure.BNode = nil
	s.Push(cur)
	for !s.IsEmpty(){
		t,_ := s.Top()
		cur = t.(*datastructure.BNode)
		if last == nil || last.Left == cur || last.Right == cur {
			if cur.Left!=nil {
				s.Push(cur.Left)
			}else if cur.Right!=nil {
				s.Push(cur.Right)
			}
		}else if cur.Left == last{
			if cur.Right!=nil {
				s.Push(cur.Right)
			}
		}else {
			fmt.Printf("%c",cur.T)
			s.Pop()
		}
		last = cur
	}
}


func Main()  {
	str := "AB##CD##EF##G#H##"
	bt := CreateBTree(str)
	bt.PreOrder()
	fmt.Println()
	bt.PreOrder2()
	bt.PreOrder3()
	bt.MorrisPreOrder()
	bt.InOrder()
	fmt.Println()
	bt.InOrder2()
	bt.MorrisInOrder()
	bt.PostOrder()
	fmt.Println()
	bt.PostOrder2()
	bt.PostOrder3()
	bt.Post()
}