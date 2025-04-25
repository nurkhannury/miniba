package main

import (
	"fmt"
	"log"
	"mini_BA/db"
	"mini_BA/handlers"
	"net/http"
)

func main() {
	db.InitDB()
	http.HandleFunc("/get-reports", handlers.GetReportsHandler)

	// Serve index.html
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Endpoint for GPT analysis
	http.HandleFunc("/analyze", handlers.AnalyzeHandler)
	http.HandleFunc("/recommend", handlers.RecommendHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🧠 Welcome to Mini Business Analyst")
}
