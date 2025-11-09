package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the Weather Dashboard home page
	http.ServeFile(w, r, "static/home.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about page
	http.ServeFile(w, r, "static/about.html")
}

func main() {
	// Weather Dashboard - A cloud-native web application
	// Built with Go, Docker, and Kubernetes
	
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)

	log.Println("Weather Dashboard starting on port 8000...")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
