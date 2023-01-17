package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	
	data := map[string]interface{} {
		"status":  true,
		"message":  "Welcome to Taocoder Go API",
		"data": nil,
	}

	res, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main()  {
	http.HandleFunc("/", index)

	// Listen
	err := http.ListenAndServe(":7000", nil)

	// Error Handler
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server is closed");
	}
	
	if err != nil {
		os.Exit(1)
	}
}

type Response struct {
	status bool
	message string
	data interface{}
}