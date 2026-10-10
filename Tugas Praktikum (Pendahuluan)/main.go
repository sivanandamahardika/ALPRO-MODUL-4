package main
import "fmt"

func main() {
	var n int
	var beli bool

	fmt.Scan(&n)
	beli = n >= 1500

	fmt.Println(beli)

}