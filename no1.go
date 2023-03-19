package main
//pokoknya komentar
import "fmt"

func main() {
	var b, faktor, nFaktor int
	fmt.Scan(&b)
	nFaktor = 0
	faktor = 1
	for faktor <= b {
		if b%faktor == 0 {
			fmt.Print(faktor, " ")
			nFaktor += 1
		}
		faktor++
	}
	fmt.Println("\nPrima:", nFaktor == 2)
}
