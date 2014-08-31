package functions

import "fmt"

func add(x int, y int) int {
    return x + y
}

func sub(x int, y int) int {
    return x - y
}

func main() {
    fmt.Println(add(2, 2))
    fmt.Println(sub(100, 10))
    
    fmt.Println(addCondense(2, 2))
    fmt.Println(subCondense(100, 10))
    
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    
    fmt.Println(split(17))
}
