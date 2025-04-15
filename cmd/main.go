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

	maxSubstr := longestSubstring(req.Input)
	c.JSON(http.StatusOK, gin.H{"result": maxSubstr})
}

func longestSubstring(s string) string {
	var l, r, maxLen int
	stringLen := len(s)
	longestSubstring := ""
	if len(s) == 0 {
		return ""
	}
	mp := make(map[string]bool)
	runes := []rune(s)

	for r < stringLen {
		if _, ok := mp[string(runes[r])]; ok {
			delete(mp, string(runes[l]))
			l++
		} else {
			mp[string(runes[r])] = true
			currentLen := r - l + 1
			if currentLen > maxLen {
				longestSubstring = string(runes[l : r+1])
				maxLen = currentLen
			}
			r++
		}
	}

	return longestSubstring
}
