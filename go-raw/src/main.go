package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/levigross/grequests"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Golang on Hasura")
}

// func sayPongJSON(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprint(w, `{"message":"pong"}`)
// }

func main() {
	http.HandleFunc("/", sayHello)
	// http.HandleFunc("/ping", sayPongJSON)

	http.HandleFunc("/get_articles", getArticles)
	http.HandleFunc("/profile", getProfile)

	// get port env var
	port := "8080"
	portEnv := os.Getenv("PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}

	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// listen and serve on 0.0.0.0:8080 by default
	// set environment variable PORT if you want to change port
}

// Endpoints demonstrating Hasura Backend Features:
func getArticles(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !resp.Ok {
		err := fmt.Errorf("code: %d, data: %s", resp.StatusCode, string(resp.Bytes()))
		log.Printf("error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"response": %s}`, string(resp.Bytes()))
	return
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	baseDomain := r.Header.Get("X-Hasura-Base-Domain")
	if len(baseDomain) == 0 {
		http.Error(w, "This URL works only on Hasura clusters", http.StatusInternalServerError)
		return
	}
	role := r.Header.Get("X-Hasura-Role")
	if role == "user" {
		userId := r.Header.Get("X-Hasura-User-Id")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"userId": %s}`, userId)
	} else {
		redirectUrl := fmt.Sprintf(
			"http://auth.%s/ui?redirect_to=http://app.%s/profile",
			baseDomain, baseDomain,
		)
		http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	}
}
