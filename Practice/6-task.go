// 6. Uch xonali son berilgan. Uni o’nliklar xonasidagi raqam bilan birliklar xonasidagi raqamni almashtirishdan xosil bo’lgan sonni toping.

package main

import (
    "fmt"
)

func main() {
    var number int
    fmt.Println("Enter number:")
    fmt.Scanln(&number)

    A := number % 10
    B := (number / 10) % 10
    C := number / 100

    result := C*100 + A*10 + B

    fmt.Println("result:", result)
}
