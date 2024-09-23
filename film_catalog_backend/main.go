package main

import (
	"film_catalog_backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database
	dsn := "host=localhost user=postgres password=admin dbname=film_catalog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	log.Println("Berhasil terhubung ke database!")

	// Inisialisasi router
	router := gin.Default()

	// Middleware untuk menyertakan database ke context Gin
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Setup routes
	routes.SetupRouter(router)

	// Jalankan server
	router.Run(":8080") // Menjalankan server di port 8080
}
