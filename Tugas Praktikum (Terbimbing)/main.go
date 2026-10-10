package main
import "fmt"

func main() {
    var h, Batas, Tekanan float64
    const p float64 = 1025.0
    const g float64 = 9.8

    fmt.Scan(&h, &Batas)
    Tekanan = p * g * h

    fmt.Println(Tekanan)
    fmt.Println(Tekanan <= Batas)
}