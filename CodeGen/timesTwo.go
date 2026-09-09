package main

import "fmt"

//this program can multiply numbers up to 4
func main() {
    var num int
    fmt.Println("Input a number: ")
    fmt.Scanf("%d", &num)

    if num == 0 {
        fmt.Println("Your Number 0 times 2 is 0")
    } else {
        if num == 1 {
            fmt.Println("Your Number 1 times 2 is 2")
        } else {
            if num == 2 {
                fmt.Println("Your Number 2 times 2 is 4")
            } else {
                if num == 3 {
                    fmt.Println("Your Number 3 times 2 is 6")
                } else {
                    if num == 4 {
                        fmt.Println("Your Number 4 times 2 is 8")
                    } else {
                        fmt.Println("Your number is too big :(")
                    }
                }
            }
        }
    }

}