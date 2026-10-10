package main
import "fmt"
func main() {

	var bilangan, d1, d2, d3 int
	fmt.Scan(&bilangan)

	d1 = bilangan / 100
	d2 = bilangan % 100 / 10
	d3 = bilangan % 100 % 10

	fmt.Println(d1 <= d2 && d2 <= d3)
}


// Program DigitTerurutMembesar
// Kamus
//     bilangan, d1, d2, d3 : integer

// Algoritma
//     input(bilangan)

//     d1 <- bilangan div 100
//     d2 <- (bilangan mod 100) div 10
//     d3 <- bilangan mod 100 mod 10

//     output(d1 <= d2 dan d2 <= d3)