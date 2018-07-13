package main

import "fmt"

func intseq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}
func main()  {
  nextInt := intseq()
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  newInts := intseq()
  fmt.Println(newInts())
}
