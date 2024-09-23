package controllers

import (
	"film_catalog_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Untuk Menambahkan Film
func CreateMovie(c *gin.Context) {
	// Ambil koneksi database dari context
	db := c.MustGet("db").(*gorm.DB)
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Simpan ke database
	result := db.Create(&movie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save movie"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// Untuk Mengambil Daftar Film
func GetMovies(c *gin.Context) {
	// Ambil koneksi database dari context
	db := c.MustGet("db").(*gorm.DB)
	var movies []models.Movie
	db.Find(&movies)
	c.JSON(http.StatusOK, movies)
}
