package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Target struct {
	url           string
	method        string
	total         int
	fail          int
	lastError     string
	lastIsSuccess bool
	lastTime      time.Time
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
				if targets[i].method == "GET" {
					resp, err = http.Get(targets[i].url)
				} else {
					resp, err = http.Post(targets[i].url, "application/x-www-form-urlencoded", nil)
				}
				targets[i].total++
				targets[i].lastTime = time.Now()
				if err == nil && resp.StatusCode < 500 {
					targets[i].lastIsSuccess = true
				} else {
					targets[i].fail++
					targets[i].lastError = err.Error()
					targets[i].lastIsSuccess = false
				}
			}
		}
	}()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		result := ""
		for i := 0; i < len(targets); i++ {
			target := targets[i]
			if target.lastIsSuccess == false {
				result += target.url + ": fail\n"
			} else {
				result += target.url + ": success\n"
			}
		}
		c.String(200, result)
	})
	r.Run("localhost:9992")
}
