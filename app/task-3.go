package app

import "math"

type Persegi struct {
	Sisi float64
}

type Tabung struct {
	Tinggi float64
}

type Lingkaran struct {
	JariJari float64
	Tabung   Tabung
}

type Kubus struct {
	Sisi    float64
	Persegi Persegi
}

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

func (l Lingkaran) Luas() float64 {
	luas := math.Pi * l.JariJari * 2
	return luas
}

func (l Lingkaran) Keliling() float64 {
	keliling := 2 * math.Pi * l.JariJari
	return keliling
}

func (l Lingkaran) Volume() float64 {
	volume := math.Pi * math.Pow(l.JariJari, 2) * l.Tabung.Tinggi
	return volume
}

func (k Kubus) Luas() float64 {
	luas := math.Pow(k.Sisi, 2)
	return luas
}

func (k Kubus) Keliling() float64 {
	keliling := k.Sisi * 12
	return keliling
}

func (k Kubus) Volume() float64 {
	volume := math.Pow(k.Sisi, 3)
	return volume
}
