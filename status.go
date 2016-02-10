package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Target struct {
	URL           string    `json:"url"`
	Method        string    `json:"method"`
	Total         int       `json:"total"`
	Fail          int       `json:"fail"`
	LastError     string    `json:"lastError"`
	LastIsSuccess bool      `json:"lastIsSuccess"`
	LastTime      time.Time `json:"lastTime"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	now := time.Now()
	targets := []Target{
		Target{"https://yorkyao.xyz/api/version", "GET", 0, 0, "", true, now},
		Target{"https://yorkyao.xyz/", "GET", 0, 0, "", true, now},
		Target{"https://doc.yorkyao.xyz/", "GET", 0, 0, "", true, now},
		Target{"https://news.yorkyao.xyz/items", "GET", 0, 0, "", true, now},
		Target{"https://robot.yorkyao.xyz/", "POST", 0, 0, "", true, now},
		Target{"https://upload.yorkyao.xyz/api/temperary", "POST", 0, 0, "", true, now},
	}

	ticker := time.NewTicker(time.Second * 60)

	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			fmt.Println(targets)
			for i := 0; i < len(targets); i++ {
				var resp *http.Response
				var err error
				if targets[i].Method == "GET" {
					resp, err = http.Get(targets[i].URL)
				} else {
					resp, err = http.Post(targets[i].URL, "application/x-www-form-urlencoded", nil)
				}
				targets[i].Total++
				targets[i].LastTime = time.Now()
				if err == nil && resp.StatusCode < 500 {
					targets[i].LastIsSuccess = true
				} else {
					targets[i].Fail++
					targets[i].LastError = err.Error()
					targets[i].LastIsSuccess = false
					defer resp.Body.Close()
				}
			}
		}
	}()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, targets)
	})
	address := "localhost:9992"
	fmt.Println("listening: " + address)
	r.Run(address)
}
