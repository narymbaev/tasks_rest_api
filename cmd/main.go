package main

import (
	"net/http"
	"regexp"
	"tasks_test_api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/rest/substr/find", findSubstringHandler)
	r.POST("/rest/email/check", checkEmailHandler)

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

	maxSubstr := utils.LongestSubstring(req.Input)
	c.JSON(http.StatusOK, gin.H{"result": maxSubstr})
}

func checkEmailHandler(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body text not found"})
	}

	bodyText := string(bodyBytes)

	re := regexp.MustCompile(`(?i)Email:\s+([a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,})`)
	matches := re.FindAllStringSubmatch(bodyText, -1)

	emails := []string{}

	for _, match := range matches {
		if len(match) > 1 {
			emails = append(emails, match[1])
		}
	}

	c.JSON(http.StatusOK, gin.H{"emails": emails})
}
