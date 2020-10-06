package main

import "fmt"

type Queue struct {
  items []int
}

func (queue *Queue) Enqueue(value int) {
  queue.items = append(queue.items, value)
}

func (queue *Queue) Dequeue() int {
  toRemove := queue.items[0]
  queue.items = queue.items[1:]
  return toRemove
}

func main() {
  myQueueStruct := Queue{}
  myQueueStruct.Enqueue(10)
  myQueueStruct.Enqueue(29)
  myQueueStruct.Enqueue(211)
  fmt.Println(myQueueStruct)
  myQueueStruct.Dequeue()
  fmt.Println(myQueueStruct)
}
