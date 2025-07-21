package handlers

import (
	"html/template"
	"net/http"
)

// WebHandler menangani request HTTP untuk antarmuka web konversi suhu
// Fungsi ini bertugas:
// 1. Memuat template HTML dari file
// 2. Menyiapkan data untuk di-render ke template
// 3. Menampilkan halaman web ke client
// Parameter:
// - w: http.ResponseWriter untuk menulis response
// - r: *http.Request yang berisi informasi request
func WebHandler(w http.ResponseWriter, r *http.Request) {
	// Memuat template HTML dari file
	// Template berada di lokasi "templates/index.html"
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// Jika gagal memuat template, kirim error 500 Internal Server Error
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Menyiapkan data untuk di-render ke template
	// Data berupa struct dengan field Title untuk judul halaman
	data := struct {
		Title string
	}{
		Title: "ThermoConvert - Konversi Suhu", // Judul halaman web
	}

	// Render template dengan data dan kirim ke client
	if err := tmpl.Execute(w, data); err != nil {
		// Jika gagal render, kirim error 500 Internal Server Error
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}