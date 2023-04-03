package main

import "fmt"

// https://refactoring.guru/design-patterns/factory-method
func main() {
	ak47, _ := getGun("AK47")
	m16, _ := getGun("M16")
	_, e := getGun("BAZOOKA")

	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("Gun: %s", ak47.getName())
	fmt.Println()
	fmt.Printf("Power: %d", ak47.getPower())
	fmt.Println()

	fmt.Printf("Gun: %s", m16.getName())
	fmt.Println()
	fmt.Printf("Power: %d", m16.getPower())
	fmt.Println()
}

func getGun(gunType string) (IGun, error) {

	switch gunType {
	case "AK47":
		return newAK47(), nil
	case "M16":
		return newM16(), nil
	}

	return nil, fmt.Errorf("No gun type found for %s", gunType)
}

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(n string) {
	g.name = n
}

func (g *Gun) setPower(p int) {
	g.power = p
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type AK47 struct {
	Gun
}

type M16 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47",
			power: 47,
		},
	}
}

func newM16() IGun {
	return &M16{
		Gun: Gun{
			name:  "M16",
			power: 1600,
		},
	}
}
