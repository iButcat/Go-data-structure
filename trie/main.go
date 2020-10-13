package main

import "fmt"


type (
  Node struct {
    Children [26]*Node
    isEnd bool
  }
  Trie struct {
    Root *Node
  }
)

func InitTrieRoot() *Trie {
  treeResult := &Trie{Root: &Node{}}
  return treeResult
}

func (trie *Trie) Insert(word string) {
  wordLength := len(word)
  currentNode := trie.Root
  for i := 0; i < wordLength; i++ {
    characterIndex := word[i] - 'a'
    if currentNode.Children[characterIndex] == nil {
      currentNode.Children[characterIndex] = &Node{}
    }
    currentNode = currentNode.Children[characterIndex]
  }
  currentNode.isEnd = true
}

func (trie *Trie) Search(word string) bool {
  wordLength := len(word)
  currentNode := trie.Root
  for i := 0; i < wordLength; i++ {
    characterIndex := word[i] - 'a'
    if currentNode.Children[characterIndex] == nil {
      return false
    }
    currentNode = currentNode.Children[characterIndex]
  }
  if currentNode.isEnd == true {
    return true
  }
  return false
}

func main() {
  trie := InitTrieRoot()
  add := []string{
    "cheval",
    "cheveux",
    "chat",
    "eclipse",
    "eclate",
  }
  for _, v := range add{
    trie.Insert(v)
  }

  fmt.Println(trie.Search("none")) // false
  fmt.Println(trie.Search("cheval")) // true
}
