package ExampleReverse

import (
    "github.com/golang/example/stringutil"
    "fmt"
)

var fibTests = []struct {
  n        int // input
  expected int // expected result
}{
  {1, 1},
  {2, 1},
  {3, 2},
  {4, 3},
  {5, 5},
  {6, 8},
  {7, 13},
}


func ExampleHello() {
        fmt.Println("hello")
        // Output: hello
}

func ExampleSalutations() {
        fmt.Println("hello, and")
        fmt.Println("goodbye")
        // Output:
        // hello, and
        // goodbye
}


func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
