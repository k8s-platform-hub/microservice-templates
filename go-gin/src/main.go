package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Gin on Hasura",
		})
	})

	// // un-comment the following lines
	//  r.GET("/ping", func(c *gin.Context) {
	//  	c.JSON(200, gin.H{
	//  		"message": "pong",
	//  	})
	//  })

	r.GET("/profile", func(c *gin.Context) {
		baseDomain := c.GetHeader("X-Hasura-Base-Domain")
		role := c.GetHeader("X-Hasura-Role")
		if role == "user" {
			userId := c.GetHeader("X-Hasura-User-Id")
			c.JSON(http.StatusOK, gin.H{
				"userId": userId,
			})
		} else {
			c.Redirect(http.StatusTemporaryRedirect,
				fmt.Sprintf(
					"http://auth.%s/ui?redirect_to=http://app.%s/profile",
					baseDomain, baseDomain,
				),
			)
		}
	})

	r.GET("/get_articles", func(c *gin.Context) {
		dataUrl := "http://data.hasura/v1/query"
		// set data url as external one if CLUSTER_NAME is set
		clusterName := os.Getenv("CLUSTER_NAME")
		if len(clusterName) > 0 {
			dataUrl = fmt.Sprintf("https://data.%s.hasura-app.io/v1/query", clusterName)
		}
		resp, err := grequests.Post(dataUrl,
			&grequests.RequestOptions{
				JSON: map[string]interface{}{
					"type": "select",
					"args": map[string]interface{}{
						"table":   "article",
						"columns": []string{"*"},
					},
				},
			},
		)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		if !resp.Ok {
			c.JSON(500, gin.H{
				"error": fmt.Errorf("code: %d, data: %s", resp.StatusCode, string(resp.Bytes())),
			})
			return
		}

		var data interface{}
		err = resp.JSON(&data)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"response": data,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 by default
	// set environment variable PORT if you want to change port
}
