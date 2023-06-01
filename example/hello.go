package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	a := 10
    fmt.Println("a:- %T and %v", a, a)
    fmt.Println(quote.Go())
    fmt.Println("add ::", add(2,3,4,5))
}

func add(values ...int) int {
    result := 0
    for i, val := range values {
        result += val
        fmt.Println("index :- ", i)
    }
    return result
}
