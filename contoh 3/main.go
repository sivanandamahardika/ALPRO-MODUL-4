package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Scan(&beratBadan, &tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("%.2f", bmi)
}

// program TinggiBeradBadan
// kamus
// 	beratBadan, tinggiBadan, bmi : real
// algoritma
// 	input(beratBadan, tinggiBadan)
// 	bmi <- beratBadan / (tinggiBadan * tinggiBadan)
// 	output(bmi)
