package api

import (
	"fmt"
	"net/http"

	"github.com/joncalhoun/qson"
	requests "github.com/rickedb/torqueprotracker/v1/api/requests"
	services "github.com/rickedb/torqueprotracker/v1/services"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming package...")
	query := r.URL.Query()
	fmt.Println(query)

	var request requests.TorqueProRequest
	err := qson.Unmarshal(&request, r.URL.RawQuery)
	if err == nil {
		fmt.Println(request)
		services.Enqueue(&request)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Finished Processing!")
	w.Write([]byte("OK!"))
}
