package main

import (
   "fmt"
   "actions/math"
)

func main() {
   fmt.Println(math.Sum(5, 4))
   fmt.Println(math.Diff(5, 4))
   fmt.Println(math.Mult(5, 4))
   fmt.Println(math.Div(20, 4))
}
