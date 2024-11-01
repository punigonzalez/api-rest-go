package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// defino estructura de producto (crear clase)
type product struct {
	ID    uint    ` gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Color string  `json:"color"`
	Price float32 `json:"price"`
}

// inicializo un slice (instanciar objetos)
/*var products = []product{
	{ID: "1", Name: "Iphone 14", Brand: "Apple", Color: "Black", Price: 450},
	{ID: "2", Name: "Iphone 15", Brand: "Apple", Color: "Midnight", Price: 650},
	{ID: "3", Name: "Iphone 16", Brand: "Apple", Color: "Red", Price: 899.99},
}*/

var (
	db       *gorm.DB
	products []product
	err      error
)

// config db
func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/api_rest_goDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// migrar el modelo Product a la base de datos
	db.AutoMigrate(&product{})
}

// funcion para obtener toodos los productos
func getProducts(c *gin.Context) {
	var products []product
	db.Find(&products)
	c.IndentedJSON(http.StatusOK, products)
}

// funcion para agregar producto
func addProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	db.Create(&newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

// funcion para obtener producto por id
func getProductById(c *gin.Context) {
	id := c.Param("id")

	var product product
	result := db.First(&product, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado."})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}

// funcion para eliminar producto por id
func deleteProductById(c *gin.Context) {
	id := c.Param("id")

	result := db.Delete(&product{}, id)
	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado."})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Producto eliminado."})
}

func main() {
	fmt.Println("********************************************")
	fmt.Println("*                                          *")
	fmt.Println("*          SERVIDOR INICIADO               *")
	fmt.Println("*                                          *")
	fmt.Println("********************************************")

	//INICIALIZO DB
	initDB()

	//rutas http
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/addProduct", addProduct)
	router.GET("/products/:id", getProductById)
	router.DELETE("/deleteProduct/:id", deleteProductById)

	router.Run("localhost:8080")
}
