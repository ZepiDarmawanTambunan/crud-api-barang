package controllers

import (
	"crud-api-barang/models"
	"crud-api-barang/validators"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBarang(c *gin.Context, db *gorm.DB) {
	var barang []models.Barang
	db.Find(&barang)
	c.JSON(http.StatusOK, barang)
}

func GetBarangByID(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var barang []models.Barang
	db.First(&barang, id)
	c.JSON(http.StatusOK, barang)
}

func CreateBarang(c *gin.Context, db *gorm.DB) {
	var input models.Barang
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input termasuk validasi gambar
	if err := validators.ValidateBarangCreate(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gambarBase64 := input.Gambar
	if gambarBase64 != "" {
		// Decode base64
		decodedImage, err := base64.StdEncoding.DecodeString(strings.Split(gambarBase64, ",")[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 image"})
			return
		}

		// Generate a unique filename for the image
		fileType := strings.Split(strings.Split(gambarBase64, ";")[0], "/")[1]
		filename := fmt.Sprintf("%d-%s.%s", time.Now().Unix(), uuid.New().String(), fileType)
		filePath := fmt.Sprintf("uploads/%s", filename)

		// Save the decoded image data to a file
		file, err := os.Create(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image file"})
			return
		}
		defer file.Close()

		_, err = file.Write(decodedImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write image data"})
			return
		}

		input.Gambar = filename
	}

	db.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateBarang(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var barang models.Barang
	if err := db.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "barang not found"})
		return
	}

	var input models.Barang
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gambarBase64 := input.Gambar
	if gambarBase64 != "" {
		// Decode base64 image
		decodedImage, err := base64.StdEncoding.DecodeString(strings.Split(input.Gambar, ",")[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 image"})
			return
		}

		// Generate a unique filename for the new image
		fileType := strings.Split(strings.Split(input.Gambar, ";")[0], "/")[1]
		newFilename := fmt.Sprintf("%d-%s.%s", time.Now().Unix(), uuid.New().String(), fileType)
		newFilePath := fmt.Sprintf("uploads/%s", newFilename)

		// Save the decoded image data to a new file
		newFile, err := os.Create(newFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new image file"})
			return
		}
		defer newFile.Close()

		_, err = newFile.Write(decodedImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write new image data"})
			return
		}

		// Delete the old image file
		if barang.Gambar != "" {
			oldFilePath := fmt.Sprintf("uploads/%s", barang.Gambar)
			if err := os.Remove(oldFilePath); err != nil {
				fmt.Println("Error deleting old image:", err)
			}
		}

		// Update barang data
		barang.Gambar = newFilename
	}

	// Update the barang in the database
	db.Save(&barang)

	c.JSON(http.StatusOK, barang)
}

func DeleteBarang(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))

	var barang models.Barang
	if err := db.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
		return
	}

	// Delete the old image file
	if barang.Gambar != "" {
		oldFilePath := fmt.Sprintf("uploads/%s", barang.Gambar)
		if err := os.Remove(oldFilePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting old image"})
			return
		}
	}

	db.Delete(&barang, id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Barang with id %v is deleted", id)})
}
