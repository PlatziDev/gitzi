package main

import (
	"net/http"
	"log"
	"github.com/PlatziDev/gitzi/gitzi"
)

func main() {
	log.Println("*****GITZI*****")
	gitzi.ReadSlackUsers()
	http.HandleFunc("/gh/", gitzi.GHWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
