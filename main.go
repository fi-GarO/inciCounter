package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "counter.db")

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Print("Database connected \n")
	}

	database.AutoMigrate(&counter{})

	DB = database
}

type counter struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type CreateCounterInput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type EditCounterInput struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func CreateCounter(c *gin.Context) {
	// Validate input
	var input CreateCounterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create counter
	counter := counter{Id: input.Id, Name: input.Name, Count: input.Count}
	DB.Create(&counter)

	c.JSON(http.StatusOK, gin.H{"data": counter})
}

func FindCounters(c *gin.Context) {
	var counters []counter
	DB.Find(&counters)
	c.JSON(http.StatusOK, gin.H{"data": counters})
}

func IncrementCounter(c *gin.Context) {
	// Get model if exist
	var newCounter counter
	if err := DB.Where("id = ?", c.Param("id")).First(&newCounter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input EditCounterInput
	input.Count = newCounter.Count + 1
	DB.Model(&newCounter).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": newCounter})
}

func DecrementCounter(c *gin.Context) {

	var newCounter counter
	if err := DB.Where("id = ?", c.Param("id")).First(&newCounter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input EditCounterInput
	if newCounter.Count < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Counter already equals 0, cannot be negative."})
		return
	}

	if newCounter.Count == 1 {
		newCounter2 := counter{Id: newCounter.Id, Name: newCounter.Name, Count: 0}
		DB.Delete(&newCounter)
		DB.Create(&newCounter2)
		c.JSON(http.StatusOK, gin.H{"data": newCounter2})
		return
	}

	input.Count = newCounter.Count - 1
	DB.Model(&newCounter).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": newCounter})
}

func DeleteCounter(c *gin.Context) {
	// Get model if exist
	var newCounter counter
	if err := DB.Where("id = ?", c.Param("id")).First(&newCounter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&newCounter)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func DeleteAllCounters(c *gin.Context) {

	var counters []counter
	DB.Delete(&counters)
	c.JSON(http.StatusOK, gin.H{"data": counters})
}

func ResetCounter(c *gin.Context) {
	var input counter
	if err := DB.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	newCounter := counter{Id: input.Id, Name: input.Name, Count: 0}
	DB.Delete(&input)
	DB.Create(&newCounter)

	c.JSON(http.StatusOK, gin.H{"data": newCounter})
}

func SetCounter(c *gin.Context) {
	count := c.Param("count")
	countNumber, _ := strconv.Atoi(count)
	if countNumber < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Counter cannot be negative number!"})
		return
	}

	var input counter
	if err := DB.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	newCounter := counter{Id: input.Id, Name: input.Name, Count: countNumber}
	DB.Delete(&input)
	DB.Create(&newCounter)

	c.JSON(http.StatusOK, gin.H{"data": newCounter})
}

func main() {
	router := gin.Default()

	fmt.Print("Starting the APP\n")
	ConnectDatabase()
	router.POST("/counters", CreateCounter)
	router.PATCH("/counters/inc/:id", IncrementCounter)
	router.PATCH("/counters/set/:id/:count", SetCounter)
	router.PATCH("/counters/dec/:id", DecrementCounter)
	router.PATCH("/counters/res/:id", ResetCounter)
	router.GET("/counters", FindCounters)
	router.DELETE("/counters/:id", DeleteCounter)
	router.DELETE("/counters/del/all", DeleteAllCounters)

	router.Run(":8080")
}
