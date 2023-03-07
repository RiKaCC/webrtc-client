/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"net/http"

	"client/cmd"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rtc", rtcHandler)
	cmd.Execute()

	http.ListenAndServe(":7890", r)
}

func rtcHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "sdfsdf")
}
