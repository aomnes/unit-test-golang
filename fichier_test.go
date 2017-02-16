package ExampleReverse

import (
    "github.com/golang/example/stringutil"
    "fmt"
)

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
