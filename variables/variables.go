package main

import "fmt"

var i int
var c, python, java bool

var x, y int = 1, 2
var b, pylon, jorba = true, false, "no!"

func main() {
    fmt.Println(i, c, python, java)
    fmt.Println(x, y, b, python, jorba)
    
    // Short variable declarations
    var n, m int = 1, 2
    k := 3
    fmt.Println(n, m, k)
    
    displayTypes()
}
