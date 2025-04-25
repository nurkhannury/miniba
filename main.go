package main

import (
	"fmt"
	"log"
	"net/http"

	"mini_BA/db"
	"mini_BA/handlers"
)

func main() {
	// 1️⃣ Initialize DB
	db.InitDB()

	// 2️⃣ Create a fresh ServeMux
	mux := http.NewServeMux()

	// 3️⃣ Serve your SPA's index.html & other static files
	//    This assumes your compiled index.html, style.css, script.js live in ./static
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// 4️⃣ Serve any other assets (icons, images, etc)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// 5️⃣ API endpoints
	mux.HandleFunc("/analyze", handlers.AnalyzeHandler)
	mux.HandleFunc("/recommend", handlers.RecommendHandler)
	mux.HandleFunc("/get_reports", handlers.GetReportsHandler)

	// 6️⃣ Startup log (so you see it before blocking)
	fmt.Println("🚀 MiniBA server listening on http://localhost:8080")

	// 7️⃣ Start serving (this call blocks)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
