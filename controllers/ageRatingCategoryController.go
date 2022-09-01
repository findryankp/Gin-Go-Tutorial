package controllers

import (
	"net/http"
	"time"

	"go-rest-api-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type movieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID uint   `json:"age_rating_category_id"`
}

// GetAllMovies godoc
// @Summary Get all movies.
// @Description Get a list of Movies.
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var movies []models.Movie
	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// CreateMovie godoc
// @Summary Create New Movie.
// @Description Creating a new Movie.
// @Tags Movie
// @Param Body body movieInput true "the body to create a new movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input movieInput
	var rating models.AgeRatingCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ageRatingCategoryID not found!"})
		return
	}

	// Create Movie
	movie := models.Movie{Title: input.Title, Year: input.Year, AgeRatingCategoryID: input.AgeRatingCategoryID}
	db.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// GetMovieById godoc
// @Summary Get Movie.
// @Description Get a Movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieById(c *gin.Context) { // Get model if exist
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// UpdateMovie godoc
// @Summary Update Movie.
// @Description Update movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Param Body body movieInput true "the body to update an movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var movie models.Movie
	var rating models.AgeRatingCategory
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input movieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ageRatingCategoryID not found!"})
		return
	}

	var updatedInput models.Movie
	updatedInput.Title = input.Title
	updatedInput.Year = input.Year
	updatedInput.AgeRatingCategoryID = input.AgeRatingCategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&movie).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// DeleteMovie godoc
// @Summary Delete one movie.
// @Description Delete a movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} map[string]boolean
// @Router /movie/{id} [delete]
func DeleteMovie(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
