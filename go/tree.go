package main

import (
	"log"
)

type node struct {
	Item  string
	Left  *node
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

//层序遍历 ： [[m] [k l] [h i j] [a b c d e f]]

*/

func (tree *bst) buildTree() {
	m := &node{Item: "m"}
	tree.root = m

	k := &node{Item: "k"}
	l := &node{Item: "l"}
	m.Left = k
	m.Right = l

	h := &node{Item: "h"}
	i := &node{Item: "i"}
	k.Left = h
	k.Right = i

	a := &node{Item: "a"}
	b := &node{Item: "b"}
	h.Left = a
	h.Right = b

	c := &node{Item: "c"}
	d := &node{Item: "d"}
	i.Left = c
	i.Right = d

	j := &node{Item: "j"}
	l.Right = j

	e := &node{Item: "e"}
	f := &node{Item: "f"}
	j.Left = e
	j.Right = f

}

// 前序遍历 (根-左-右)
func (tree *bst) inOrder() {
	var inner func(n *node)
	inner = func(n *node) {
		if n == nil {
			return
		}
		log.Println(n.Item)
		inner(n.Left)
		inner(n.Right)
	}
	inner(tree.root)
}

// 中序遍历 （左-根-右）
func (tree *bst) midOrder() {
	var inner func(n *node)
	inner = func(n *node) {
		if n == nil {
			return
		}
		inner(n.Left)
		log.Println(n.Item)
		inner(n.Right)
	}
	inner(tree.root)
}

// 后序遍历 （左-右-根）
func (tree *bst) lastOrder() {
	var inner func(n *node)
	inner = func(n *node) {
		if n == nil {
			return
		}
		inner(n.Left)
		inner(n.Right)
		log.Println(n.Item)
	}
	inner(tree.root)
}

// 层序遍历
func (tree *bst) levelOrder() {
	var inner func(n *node, level int)
	res := make([][]string, 0)
	inner = func(n *node, level int) {
		if n == nil {
			return
		}
		if level == len(res) {
			res = append(res, []string{})
		}
		res[level] = append(res[level], n.Item)
		inner(n.Left, level+1)
		inner(n.Right, level+1)
	}
	inner(tree.root, 0)
	log.Println("levelOrder:", res)
}

// 前序遍历 迭代 (栈)（根-左-右）
func (tree *bst) inOrderStack() {
	stack := []*node{}
	node := tree.root
	res := []string{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Item)
			stack = append(stack, node)
			node = node.Left
		}
		top := len(stack) - 1   // 栈顶
		node = stack[top].Right // 出栈
		stack = stack[:top]
	}
	log.Println("inOrderStack:", res)
}

// 中序遍历 迭代 (栈) （左-根-右）
func (tree *bst) midOrderStack() {
	stack := []*node{}
	node := tree.root
	res := []string{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		top := len(stack) - 1 // 栈顶
		node = stack[top]     // 出栈
		res = append(res, node.Item)
		stack = stack[:top]
		node = node.Right
	}
	log.Println("midOrderStack:", res)
}

// 后序遍历  迭代 (栈)（左-右-根）
func (tree *bst) lastOrderStack() {
	var prev *node
	stack := []*node{}
	res := []string{}
	node := tree.root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		top := len(stack) - 1
		node = stack[top]
		stack = stack[:top]
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Item)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	log.Println("lastOrderStack:", res)
}

// 层序遍历 迭代
func (tree *bst) levelOrderStack() {
	res := make([][]string, 0)
	if tree.root == nil {
		return
	}
	queue := []*node{tree.root}
	for level := 0; 0 < len(queue); level++ {
		res = append(res, []string{})
		next := []*node{}
		for j := 0; j < len(queue); j++ {
			n := queue[j]
			res[level] = append(res[level], n.Item)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		queue = next
	}
	log.Println("levelOrderStack:", res)
}

func main() {
	tree := &bst{}
	tree.buildTree()

	//tree.inOrder()
	//tree.midOrder()
	//tree.lastOrder()
	//tree.levelOrder()
	//tree.inOrderStack()
	//tree.midOrderStack()
	//tree.lastOrderStack()
	tree.levelOrderStack()
}
