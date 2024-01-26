package app

import "fmt"

type DeretBilangan struct {
	Limit int
}

func (d *DeretBilangan) Prima() {
	fmt.Printf("Deret bilangan prima dengan limit %d: ", d.Limit)

	for i := 1; i <= d.Limit; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d", i)
			if i < d.Limit {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}

func (d *DeretBilangan) Ganjil() {
	fmt.Printf("Deret bilangan ganjil dengan limit %d: ", d.Limit)

	for i := 1; i <= d.Limit; i += 2 {
		fmt.Printf("%d", i)
		if i < d.Limit-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func (d *DeretBilangan) Genap() {
	fmt.Printf("Deret bilangan genap dengan limit %d: ", d.Limit)

	for i := 2; i <= d.Limit; i += 2 {
		fmt.Printf("%d", i)
		if i < d.Limit {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func (d *DeretBilangan) Fibonacci() {
	fmt.Printf("Deret Fibonacci hingga %d: ", d.Limit)

	a, b := 0, 1
	for a <= d.Limit {
		fmt.Printf("%d", a)
		if a < d.Limit-1 {
			fmt.Print(" ")
		}
		a, b = b, a+b
	}
	fmt.Println()
}
