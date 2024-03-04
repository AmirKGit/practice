// 2.n musbat soni berilgan, 2 va n   ga karrali eng kichik musbat sonni qaytaring. 

package main

import "fmt"

func main() {
    var n int
    fmt.Println("Enter number")
    fmt.Scanln(&n)

    var a int
    if n%2 == 0 {
        a = 2 * n
    } else {
        a = 2 * n
    }

    fmt.Printf("result:", n, a)
}
