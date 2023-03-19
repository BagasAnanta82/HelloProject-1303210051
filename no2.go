package main

import "fmt"

func main() {
	var gram, kg, gr, harga_kg, harga_gr, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	gr = gram % 1000
	fmt.Println("Detail berat", kg, "kg +", gr, "gr")

	harga_kg = 10000 * kg
	if gr >= 500 {
		harga_gr = gr * 5
	} else {
		harga_gr = gr * 15
	}
	fmt.Println("Detail biaya : Rp", harga_kg, "+ Rp", harga_gr)

	if kg > 10 {
		total = harga_kg
	} else {
		total = harga_kg + harga_gr
	}
	fmt.Println("Total biaya : Rp", total)
}
