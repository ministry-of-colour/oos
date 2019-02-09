package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/husobee/vestigo"
	api "github.com/ministry-of-colour/theoldowlsscarf.com/api"
	"github.com/sirupsen/logrus"
)

func main() {
	port := 8000
	flag.IntVar(&port, "port", 8000, "port to run the api-server from")

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	api := api.NewHTTP(log)

	// Setup the router
	router := vestigo.NewRouter()
	vestigo.AllowTrace = true

	// Setting up router global  CORS policy
	// These policy guidelines are overriddable at a per resource level shown below
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "test.com"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-Header", "X-Y-Header"},
		MaxAge:           3600 * time.Second,
		AllowHeaders:     []string{"X-Header", "X-Y-Header"},
	})

	router.Get("/", api.Default)
	router.Get("/hello", api.Hello)
	router.Get("/stock/:brand", api.StockReport)

	// Below Applies Local CORS capabilities per Resource (both methods covered)
	// by default this will merge the "GlobalCors" settings with the resource
	// cors settings.  Without specifying the AllowMethods, the router will
	// accept any Request-Methods that have valid handlers associated
	router.SetCors("/stock/:brand", &vestigo.CorsAccessControl{
		AllowMethods: []string{"GET"},                    // only allow cors for this resource on GET calls
		AllowHeaders: []string{"X-Header", "X-Z-Header"}, // Allow this one header for this resource
	})

	log.Println("Starting API server", api.Version(), " Port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
