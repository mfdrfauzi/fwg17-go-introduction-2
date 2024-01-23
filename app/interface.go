package app

import "math"

type hitung2d interface {
	Luas() float64
	Keliling() float64
}

type hitung3d interface {
	Volume() float64
}

type Hitung interface {
	hitung2d
	hitung3d
}

type Lingkaran struct {
	JariJari float64
	Tinggi   float64
}

func (l *Lingkaran) Luas() float64 {
	luas := math.Pi * l.JariJari * 2
	return luas
}

func (l *Lingkaran) Keliling() float64 {
	keliling := 2 * math.Pi * l.JariJari
	return keliling
}

func (l *Lingkaran) Volume() float64 {
	volume := math.Pi * math.Pow(l.JariJari, 2) * l.Tinggi
	return volume
}
