package main

import (
	"log"
	"net/http"
	"os"

	apiv1 "github.com/rickedb/torqueprotracker/v1/api"
	services "github.com/rickedb/torqueprotracker/v1/services"
	"github.com/robfig/cron/v3"
)

func main() {
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")
	c := cron.New()
	c.AddFunc("* * * * *", func() { services.Save(database) })
	c.Start()
	http.HandleFunc("/upload", apiv1.Upload)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
