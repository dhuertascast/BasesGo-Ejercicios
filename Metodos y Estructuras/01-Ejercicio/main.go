package main

import (
	"fmt"
)

// Product es una estructura que representa un producto.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Products es un slice global de Product instanciado con valores.
var Products = []Product{
	{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Descripción 1", Category: "Categoría 1"},
	{ID: 2, Name: "Producto 2", Price: 29.99, Description: "Descripción 2", Category: "Categoría 2"},
	{ID: 3, Name: "Producto 3", Price: 39.99, Description: "Descripción 3", Category: "Categoría 3"},
}

// Save añade el producto al slice Products.
func (p *Product) Save() {
	Products = append(Products, *p)
}

// GetAll imprime todos los productos guardados en Products.
func GetAll() {
	fmt.Println("Lista de productos:")
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
	fmt.Println()
}

// getById retorna el producto correspondiente al ID pasado como parámetro.
func getById(id int) (Product, error) {
	for _, product := range Products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("Producto con ID %d no encontrado", id)
}

func main() {
	// Crear un nuevo producto
	newProduct := Product{ID: 4, Name: "Nuevo Producto", Price: 49.99, Description: "Descripción del nuevo producto", Category: "Categoría 4"}

	// Guardar el nuevo producto
	newProduct.Save()

	// Imprimir todos los productos
	GetAll()

	// Obtener un producto por ID
	idToGet := 2
	product, err := getById(idToGet)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Producto con ID %d encontrado: %v\n", idToGet, product)
	}
}
