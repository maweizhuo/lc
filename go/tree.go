package main

import "log"

type node struct {
	Item string
	Left *node
	Right *node
}

type bst struct {
	root *node
}

/*
树的层级遍历
        m
     k     l
  h    i     j
a  b  c  d  e  f


//先序遍历（根左右）：m k h a b i c d l j e f

//中序遍历（左根右）：a h b k c i d m l e j f

//后序遍历（左右根）：a b h c d i k e f j l m

*/

func (tree *bst) buildTree(){
	m:=&node{Item:"m"}
	tree.root=m

	k:=&node{Item:"k"}
	l:=&node{Item:"l"}
	m.Left=k
	m.Right=l

	h:=&node{Item:"h"}
	i:=&node{Item:"i"}
	k.Left=h
	k.Right=i

	a:=&node{Item:"a"}
	b:=&node{Item:"b"}
	h.Left=a
	h.Right=b

	c:=&node{Item:"c"}
	d:=&node{Item:"d"}
	i.Left=c
	i.Right=d

	j:=&node{Item:"j"}
	l.Right=j

	e:=&node{Item:"e"}
	f:=&node{Item:"f"}
	j.Left=e
	j.Right=f

}

// 前序遍历 (根-左-右)
func (tree *bst)inOrder()  {
	var inner func(n *node)
	  inner= func(n *node) {
		if n==nil{
			return
		}
		log.Println(n.Item)
		inner(n.Left)
		inner(n.Right)
	}
	inner(tree.root)
}

// 中序遍历 （左-根-右）
func (tree *bst) midOrder(){
	var inner func(n *node)
	inner= func(n *node) {
		if n==nil{
			return
		}
		inner(n.Left)
		log.Println(n.Item)
		inner(n.Right)
	}
	inner(tree.root)
}

// 后序遍历 （左-右-根）
func (tree *bst)lastOrder()  {
	var inner func(n *node)
	inner= func(n *node) {
		if n==nil{
			return
		}
		inner(n.Left)
		inner(n.Right)
		log.Println(n.Item)
	}
	inner(tree.root)
}

func main()  {
	 tree:=&bst{}
	 tree.buildTree()

	 //tree.inOrder()
	 //tree.midOrder()
	 tree.lastOrder()
}