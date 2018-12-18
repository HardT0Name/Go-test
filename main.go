package main

import "fmt"

func main() {
    var i = 3
    var j = 20
    movij( &i , &j )
    fmt.Println(i,j)
}

func movij( i , j ...*int) {
    var mid int
    mid = *i
    *i = *j
    *j = mid
}