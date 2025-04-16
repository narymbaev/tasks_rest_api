package main

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"tasks_test_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb *redis.Client
var ctx = context.Background()

func main() {
	r := gin.Default()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r.POST("/rest/substr/find", findSubstringHandler)
	r.POST("/rest/email/check", checkEmailHandler)
	r.POST("/rest/iin/check", checkIINHandler)
	r.POST("/rest/counter/add/:i", addCounterHandler)
	r.POST("/rest/counter/sub/:i", subCounterHandler)
	r.GET("/rest/counter/val", getCounterValueHandler)

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
		return
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

func checkIINHandler(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse raw data"})
		return
	}

	bodyText := string(bodyBytes)

	re := regexp.MustCompile(`(?i)IIN:\s*(\d{12})`)
	matches := re.FindAllStringSubmatch(bodyText, -1)

	iins := []string{}

	for _, match := range matches {
		if len(match) > 1 {
			iins = append(iins, match[1])
		}
	}

	c.JSON(http.StatusOK, gin.H{"iins": iins})
}

func addCounterHandler(c *gin.Context) {
	param := c.Param("i")
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse param 'i'"})
		return
	}

	res, err := rdb.IncrBy("counter", val).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not increment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "counter incremented, current value: " + strconv.FormatInt(res, 10)})
}

func subCounterHandler(c *gin.Context) {
	param := c.Param("i")
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse param 'i'"})
		return
	}

	res, err := rdb.DecrBy("counter", val).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not decrement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "counter decremented, current value: " + strconv.FormatInt(res, 10)})
}

func getCounterValueHandler(c *gin.Context) {
	counterValue, err := rdb.Get("counter").Result()
	if err != nil {
		// Set to 0
		err = rdb.Set("counter", 0, 0).Err()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not get counter IN"})
			return
		}

		// Get value
		counterValue, err = rdb.Get("counter").Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not get counter IN"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"value": counterValue})
}
