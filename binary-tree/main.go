package main

import (
  "fmt"
)

type Node struct {
  Key int
  Left *Node
  Right *Node
}

func (n *Node) Insert(key int) {
  if n.Key < key {
    if n.Right == nil {
      n.Right = &Node{Key: key}
    } else {
      n.Right.Insert(key)
    }
  } else if n.Key > key {
    if n.Left == nil {
      n.Left = &Node{Key: key}
    } else {
      n.Left.Insert(key)
    }
  }
}

func (n *Node) Search (key int) bool {
  if n == nil {
    return false
  }
  if n.Key < key {
    return n.Right.Search(key)
  } else if n.Key > key {
    return n.Left.Search(key)
  }
  return true
}


func main() {
  myTree := Node{}

  myTree.Insert(5)
  myTree.Insert(17)
  myTree.Insert(87)
  myTree.Insert(34)
  myTree.Insert(189)
  myTree.Insert(300)
  myTree.Insert(123)

  fmt.Println(myTree.Search(34))
  fmt.Println(myTree.Search(2))
}
