package main

import "fmt"

type Stack struct {
  items []int
}

func (stack *Stack) push(value int) {
  stack.items = append(stack.items, value)
}

func (stack *Stack) pop() int {
  length := len(stack.items)-1
  toRemove := stack.items[length]
  stack.items = stack.items[:length]
  return toRemove
}

func main() {
  myStackStruct := Stack{}
  myStackStruct.push(30)
  myStackStruct.push(60)
  myStackStruct.push(22)
  fmt.Println(myStackStruct)
  myStackStruct.pop()
  fmt.Println(myStackStruct)
}
