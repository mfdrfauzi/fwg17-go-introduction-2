package main

import (
	"fmt"

	"gituhub.com/mfdrfauzi/fwg17-go-introduction-2/app"
)

func main() {
	//---Rounding---
	fmt.Println("---Rounding(Pembulatan)---")
	var n float64

	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)

	app.Rounding(n)

	//---Deret---
	fmt.Println("\n---Deret Bilangan---")
	deret := app.DeretBilangan{Limit: 20}
	deret.Prima()
	deret.Ganjil()
	deret.Genap()
	deret.Fibonacci()

	//Interface
	fmt.Println("\n---Interface---")
	var lingkaran = &app.Lingkaran{
		JariJari: 5.0,
		Tinggi:   3.0,
	}

	fmt.Printf("Luas Lingkaran: %.2f\n", lingkaran.Luas())
	fmt.Printf("Keliling Lingkaran: %.2f\n", lingkaran.Keliling())
	fmt.Printf("Volume Lingkaran: %.2f\n", lingkaran.Volume())
}
