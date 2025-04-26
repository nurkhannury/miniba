package main

import (
	"log"
	"net/http"
	"os"

	"mini_BA/db"
	"mini_BA/handlers"
)

func main() {
	// 1️⃣ Initialize DB
	db.InitDB()

	// 2️⃣ Create a new ServeMux
	mux := http.NewServeMux()

	// 3️⃣ Serve static files (index.html, CSS, JS) from ./static
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// 4️⃣ Serve other assets (images, favicon, etc) from ./assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// 5️⃣ API endpoints
	mux.HandleFunc("/analyze", handlers.AnalyzeHandler)
	mux.HandleFunc("/recommend", handlers.RecommendHandler)
	mux.HandleFunc("/get_reports", handlers.GetReportsHandler)

	// 6️⃣ Pick up the port from environment (Render, Heroku etc.)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 7️⃣ Log & start server
	log.Printf("🚀 MiniBA server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
