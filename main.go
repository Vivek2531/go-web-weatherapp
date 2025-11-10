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

func rootRedirect(w http.ResponseWriter, r *http.Request) {
	// Redirect root path to /home
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		return
	}
	// Handle 404 for other paths
	http.NotFound(w, r)
}

func main() {
	// Weather Dashboard - A cloud-native web application
	// Built with Go, Docker, and Kubernetes
	
	http.HandleFunc("/", rootRedirect)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/about", aboutPage)

	log.Println("Weather Dashboard starting on port 8000...")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
