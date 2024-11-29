package main

import (
	"log"
	"net/http"

	"simple-file-manager/config"
	"simple-file-manager/handlers"
)

func main() {

	// Load the configuration
	config.LoadConfig()

	// Serve static files (CSS & JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register handlers
	http.HandleFunc("/", handlers.ListFilesHandler)
	http.HandleFunc("/download-all", handlers.DownloadAllHandler)
	http.HandleFunc("/delete", handlers.DeleteFileHandler)
	http.HandleFunc("/rename", handlers.RenameFileHandler)

	// Start the server

	log.Printf("Server is running on http://localhost%s\n", config.PORT)
	log.Fatal(http.ListenAndServe(config.PORT, nil))
}
