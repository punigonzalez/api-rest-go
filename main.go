package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tipo de dato *objeto?*
type product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Color string  `json:"color"`
	Price float32 `json:"price"`
}

// instancio los objetos producto
var products = []product{
	{ID: "1", Name: "Iphone 14", Brand: "Apple", Color: "Black", Price: 450},
	{ID: "2", Name: "Iphone 15", Brand: "Apple", Color: "Midnight", Price: 650},
	{ID: "3", Name: "Iphone 16", Brand: "Apple", Color: "Red", Price: 899.99},
}

// funcion para obtener toodos los productos
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// funcion para agregar producto
func postProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)

	c.IndentedJSON(http.StatusCreated, products)
}

// funcion para obtener producto por id
func getProductById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado."})
}

// funcion para eliminar producto por id
func deleteProductById(c *gin.Context) {
	id := c.Param("id")

	for index, a := range products {
		if a.ID == id {
			// Eliminar el producto de la lista
			products = append(products[:index], products[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Producto eliminado."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado."})
}

func main() {
	fmt.Println("********************************************")
	fmt.Println("*                                          *")
	fmt.Println("*          SERVIDOR INICIADO               *")
	fmt.Println("*                                          *")
	fmt.Println("********************************************")

	//rutas http
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/addProduct", postProduct)
	router.GET("/products/:id", getProductById)
	router.DELETE("/deleteProduct/:id", deleteProductById)

	router.Run("localhost:8080")
}
