package main

import "fmt"

type car interface {
	Sign() string
	accept(visitor)
}
type Toyota struct {
	price int
}

func (t *Toyota) accept(v visitor) {
	v.visitForToyota(t)
}

func (t*Toyota) Sign() string {
	return "Toyota"
}
type Hyundai struct {
	price int
}

func (h *Hyundai) accept(v visitor) {
	v.visitForHyundai(h)
}

func (h*Hyundai) Sign() string {
	return "Hyundai"
}
type Mers struct {
	price int
}

func (m* Mers) accept(v visitor) {
	v.visitForMers(m)
}

func (m*Mers) Sign() string {
	return "rectangle"
}
type visitor interface {
	visitForToyota(*Toyota)
	visitForHyundai(*Hyundai)
	visitForMers(*Mers)
}
type CalculateForMonth struct {
	month int
}

func (c*CalculateForMonth) visitForToyota(t *Toyota) {
	fmt.Println("calculate for months")
}

func (c*CalculateForMonth) visitForHyundai(h*Hyundai) {
	fmt.Println("calculate for months")
}
func (c* CalculateForMonth) visitForMers(m*Mers) {
	fmt.Println("calculate for months")
}
type calculateForYear struct {
	year int
}

func (c*calculateForYear) visitForToyota(t*Toyota) {
	fmt.Println("Calculating for year")
}

func (c*calculateForYear) visitForHyundai(h*Hyundai) {
	fmt.Println("Calculating for year")
}
func (c*calculateForYear) visitForMers(m*Mers) {
	fmt.Println("Calculating for year")
}

func main() {
	toyota:=&Toyota{price: 12000000}
	hyundai:=&Hyundai{price: 8000000}
	Mers:=&Mers{price: 10000000}
	calForMonth:=&CalculateForMonth{}
	toyota.accept(calForMonth)
	hyundai.accept(calForMonth)
	Mers.accept(calForMonth)

	calForYear:=&calculateForYear{}
	toyota.accept(calForYear)
	hyundai.accept(calForYear)
	Mers.accept(calForYear)



}