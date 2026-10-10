package main

import "fmt"

func main() {
    var totalBelanja, diskon int
    fmt.Scan(&totalBelanja, &diskon)
    hargaAkhir := totalBelanja - (totalBelanja * diskon / 100)
    fmt.Println(hargaAkhir)
}