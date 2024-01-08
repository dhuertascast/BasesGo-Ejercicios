package main

import "fmt"

// Product es una interfaz que define el método Price.
type Product interface {
	Price() float64
}

// AdditionalCost es una interfaz que define el método AdditionalCost.
type AdditionalCost interface {
	AdditionalCost() float64
}

// BaseProduct es una estructura base que contiene el precio del producto.
type BaseProduct struct {
	Price float64
}

// SmallProduct es una estructura que compone BaseProduct y implementa la interfaz Product.
type SmallProduct struct {
	BaseProduct
}

// Price devuelve el precio total para productos de tipo Small.
func (p *SmallProduct) Price() float64 {
	return p.BaseProduct.Price
}

// MediumProduct es una estructura que compone BaseProduct e implementa la interfaz Product y AdditionalCost.
type MediumProduct struct {
	BaseProduct
}

// AdditionalCost devuelve el costo adicional para productos de tipo Medium.
func (p *MediumProduct) AdditionalCost() float64 {
	return p.BaseProduct.Price * 0.03
}

// Price devuelve el precio total para productos de tipo Medium.
func (p *MediumProduct) Price() float64 {
	return p.BaseProduct.Price + p.AdditionalCost()
}

// LargeProduct es una estructura que compone BaseProduct e implementa la interfaz Product y AdditionalCost.
type LargeProduct struct {
	BaseProduct
}

// AdditionalCost devuelve el costo adicional para productos de tipo Large.
func (p *LargeProduct) AdditionalCost() float64 {
	return p.BaseProduct.Price*0.06 + 2500
}

// Price devuelve el precio total para productos de tipo Large.
func (p *LargeProduct) Price() float64 {
	return p.BaseProduct.Price + p.AdditionalCost()
}

// ProductFactory es una función que retorna un producto según el tipo especificado.
func ProductFactory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return &SmallProduct{BaseProduct: BaseProduct{Price: price}}
	case "Medium":
		return &MediumProduct{BaseProduct: BaseProduct{Price: price}}
	case "Large":
		return &LargeProduct{BaseProduct: BaseProduct{Price: price}}
	default:
		return nil
	}
}

func main() {
	// Crear productos usando la factory
	smallProduct := ProductFactory("Small", 100)
	mediumProduct := ProductFactory("Medium", 150)
	largeProduct := ProductFactory("Large", 200)

	// Imprimir los precios totales
	fmt.Printf("Total Price (Small): $%.2f\n", smallProduct.Price())
	fmt.Printf("Total Price (Medium): $%.2f\n", mediumProduct.Price())
	fmt.Printf("Total Price (Large): $%.2f\n", largeProduct.Price())
}
