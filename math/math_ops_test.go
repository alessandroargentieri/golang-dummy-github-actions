package math

import (
    "testing"
)

func TestSum(t *testing.T) {
   a := 5
   b := 4
   expected := 9
   actual := Sum(a, b)
   if actual != expected {
      t.Errorf("TestSum failed! Expected %d, got %d", expected, actual)
   }
}
