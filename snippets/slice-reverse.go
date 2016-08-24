package main

import "fmt"

func ReverseIndex(s []int) []int {
    r := make([]int, len(s))
    for i := len(s); i > 0; i-- {
        r[len(s) - i] = s[i - 1]
    }
    return r
}
func Reverse(slice []int) []int {
    length := len(slice)
    reversed := make([]int, length)
    for i := range slice {
        reversed[i] = slice[length - i - 1]
    }
    return reversed
}


func main() {
    s := []int{5, 4, 3, 2, 1}
    fmt.Println(Reverse(s))
    // fmt.Println(ReverseIndex(s))

}
