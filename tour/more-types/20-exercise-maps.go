package main

import (
  "fmt"
  "strings"
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  var words = strings.Fields(s)
  var count = make(map[string]int)
  for _, w := range words {
    count[w] += 1
  }
  fmt.Println(count)
  return count
}

func main() {
  wc.Test(WordCount)
}

