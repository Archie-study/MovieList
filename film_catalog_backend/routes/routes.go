package routes

import (
	"film_catalog_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	// Route untuk menambahkan film
	router.POST("/movies", controllers.CreateMovie)

	// Route untuk mendapatkan daftar film
	router.GET("/movies", controllers.GetMovies)
}
