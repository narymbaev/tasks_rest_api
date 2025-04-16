package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/rest/substr/find", findSubstringHandler)

	r.Run(":8088")
}

func findSubstringHandler(c *gin.Context) {
	var req struct {
		Input string `json:"input"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	maxSubstr := (req.Input)
	c.JSON(http.StatusOK, gin.H{"result": maxSubstr})
}
