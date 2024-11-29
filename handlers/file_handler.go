package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"simple-file-manager/config"
	"simple-file-manager/services"
	"simple-file-manager/utils"
	"text/template"
)

// ListFilesHandler renders the file list page
func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := services.ListFiles(config.FileDirectory)
	if err != nil {
		http.Error(w, "Error reading files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Files": files,
	})
}

// DownloadAllHandler creates and serves a ZIP file of all files
func DownloadAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	zipPath := "./downloads.zip"
	err := utils.CreateZipFile(config.FileDirectory, zipPath)
	if err != nil {
		http.Error(w, "Error creating ZIP file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="downloads.zip"`)
	http.ServeFile(w, r, zipPath)
}

// DeleteFileHandler handles file deletion requests
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		File string `json:"file"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(config.FileDirectory, request.File)
	if err := os.Remove(filePath); err != nil {
		http.Error(w, "Error deleting file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RenameFileHandler handles file renaming requests
func RenameFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		OldName string `json:"oldName"`
		NewName string `json:"newName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	oldPath := filepath.Join(config.FileDirectory, request.OldName)
	newPath := filepath.Join(config.FileDirectory, request.NewName)

	// Log paths to terminal
	fmt.Println("Old Path:", oldPath)
	fmt.Println("New Path:", newPath)

	if err := os.Rename(oldPath, newPath); err != nil {
		http.Error(w, "Error renaming file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond back to the browser with the old and new paths
	response := struct {
		OldPath string `json:"oldPath"`
		NewPath string `json:"newPath"`
	}{
		OldPath: oldPath,
		NewPath: newPath,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
