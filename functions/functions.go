package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func sub(x int, y int) int {
    return x - y
}

func addCondense(x, y int) int {
    return x + y
}

func subCondense(x, y int) int {
    return x - y
}

func main() {
    fmt.Println(add(2, 2))
    fmt.Println(sub(100, 10))
    
    fmt.Println(addCondense(2, 2))
    fmt.Println(subCondense(100, 10))
}
