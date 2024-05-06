package main

import "fmt"

func main() {
    // Original slice
    slice := []int{1, 2, 3, 4, 5}

    // Index of the element to be removed
    index := 2

    // Remove the element at index
    slice = append(slice[:index], slice[index+1:]...)

    fmt.Println(slice) // Output: [1 2 4 5]
}