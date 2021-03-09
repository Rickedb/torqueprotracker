package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/joncalhoun/qson"
	requests "github.com/rickedb/torqueprotracker/v1/api/requests"
	services "github.com/rickedb/torqueprotracker/v1/services"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming package...")
	query := r.URL.Query()
	fmt.Println(query)

	var request requests.TorqueProRequest
	raw := strings.ReplaceAll(r.URL.RawQuery, "Infinity", "0")
	err := qson.Unmarshal(&request, raw)
	if err == nil {
		fmt.Println(request)
		services.Enqueue(&request)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Finished Processing!")
	w.Write([]byte("OK!"))
}
