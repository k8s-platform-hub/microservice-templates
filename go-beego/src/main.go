package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/levigross/grequests"
)

func sayHello(ctx *context.Context) {
	ctx.Output.Body([]byte("Hello from Beego on Hasura"))
}

// func sayPongJSON(ctx *context.Context) {
// 	ctx.Output.JSON(map[string]string{"message": "pong"}, true, true)
// }

func main() {
	beego.Get("/", sayHello)
	// beego.Get("/ping", sayPongJSON)

	beego.Get("/get_articles", getArticles)
	beego.Get("/profile", getProfile)

	// get port env var
	port := "8080"
	portEnv := os.Getenv("PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}

	beego.Run(fmt.Sprintf("0.0.0.0:%s", port)) // listen and serve on 0.0.0.0:8080 by default
	// set environment variable PORT if you want to change port
}

// Endpoints demonstrating Hasura Backend Features:
func getArticles(ctx *context.Context) {
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
		log.Printf("error: %s", err)
		ctx.Abort(http.StatusInternalServerError, err.Error())
		return
	}
	if !resp.Ok {
		err := fmt.Errorf("code: %d, data: %s", resp.StatusCode, string(resp.Bytes()))
		log.Printf("error: %s", err)
		ctx.Abort(http.StatusInternalServerError, err.Error())
		return
	}

	var data interface{}
	err = resp.JSON(&data)
	if err != nil {
		log.Printf("error: %s", err)
		ctx.Abort(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Output.JSON(map[string]interface{}{"response": data}, true, true)
	return
}

func getProfile(ctx *context.Context) {
	baseDomain := ctx.Input.Header("X-Hasura-Base-Domain")
	if len(baseDomain) == 0 {
		ctx.Abort(http.StatusInternalServerError, "This URL works only on Hasura clusters")
		return
	}
	role := ctx.Input.Header("X-Hasura-Role")
	if role == "user" {
		userId := ctx.Input.Header("X-Hasura-User-Id")
		ctx.Output.JSON(map[string]string{"userId": userId}, true, true)
	} else {
		redirectUrl := fmt.Sprintf(
			"http://auth.%s/ui?redirect_to=http://app.%s/profile",
			baseDomain, baseDomain,
		)
		ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
	}
}
