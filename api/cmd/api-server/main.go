package main

import (
	"flag"
	"fmt"
	"net/http"

	api "github.com/ministry-of-colour/theoldowlsscarf.com/api"
	"github.com/sirupsen/logrus"
)

func main() {
	port := 8000
	flag.IntVar(&port, "port", 8000, "port to run the api-server from")

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	api := api.NewHttp(log)
	log.Println("Starting API server", api.Version(), " Port:", port)

	http.HandleFunc("/", api.Default)
	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/stock", api.Stock)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
