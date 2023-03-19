package main

import "fmt"

func pangkat(a, b int) int {
	jumlah := 1
	for i := 0; i < b; i++ {
		jumlah *= a
	}
	return jumlah
}

func konversi(biner int, desimal *int) {
	var a, i int
	*desimal = 0
	i = 0
	for biner > 0 {
		a = biner % 10
		*desimal += a * pangkat(2, i)
		i++
		biner = biner / 10
	}
}

func main() {
	var x, y int
	fmt.Scan(&x)
	konversi(x, &y)
	fmt.Println(y)
}
