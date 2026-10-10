package main
import "fmt"

func main() {
var detik, jam, menit int
fmt.Scan(&detik)

jam = detik / 3600
menit = (detik % 3600) / 60
detik = detik % 60
fmt.Printf("%d jam %d menit %d detik\n", jam, menit, detik)
}


// Program KonversiDetik
// Kamus
//     detik, jam, menit : integer
// Algoritma
//     input(detik)
//     jam <- detik div 3600
//     menit <- (detik mod 3600) div 60
//     detik <- detik mod 60
//     output(jam, " jam ", menit, " menit dan ", detik, " detik")