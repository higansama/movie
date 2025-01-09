package utils

import (
	"github.com/gin-gonic/gin"
)

// RespondWithJSON is a utility function to respond with JSON data.
func RespondWithJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// RespondWithError is a utility function to respond with an error message.
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

// HashString is a utility function to hash a string (for example, passwords).
// You can implement your hashing logic here.
func HashString(input string) string {
	// Implement hashing logic
	return input // Placeholder
}
