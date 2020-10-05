package main

import "fmt"

type (
  node struct {
    data int
    next *node
  }

  linkedList struct {
    head *node
    length int
  }
)

func (l *linkedList) append(n *node) {
  second := l.head
  l.head = n
  l.head.next = second
  l.length++
  }

func (l *linkedList) delete(value int) {
  if l.length == 0 {
    return
  }
  if l.head.data == value {
    l.head = l.head.next
    l.length--
    return
  }
  valueToBeDelete := l.head
  for valueToBeDelete.next.data != value {
    if valueToBeDelete.next.next == nil {
      return
    }
    valueToBeDelete = valueToBeDelete.next
  }
  valueToBeDelete.next = valueToBeDelete.next.next
  l.length--
}

func (l linkedList) printLinkedList() {
  valueToBePrint := l.head
  for l.length != 0 {
    fmt.Println("%d ", valueToBePrint.data)
    valueToBePrint = valueToBePrint.next
    l.length--
  }
  fmt.Println("\n")
}

func main() {
  myLinkedList := linkedList{}
  firstNode := &node{data:12}
  secondNode := &node{data:38}
  thirdNode := &node{data:42}
  fourNode := &node{data:88}
  fiveNode := &node{data: 91}
  myLinkedList.append(firstNode)
  myLinkedList.append(secondNode)
  myLinkedList.append(thirdNode)
  myLinkedList.append(fourNode)
  myLinkedList.append(fiveNode)
  myLinkedList.printLinkedList()
  myLinkedList.delete(88)
  myLinkedList.printLinkedList()
}
