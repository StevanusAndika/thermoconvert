package main

import (
	"log"
	"net/http"
	"thermoconvert/handlers"  // Import package handlers lokal

	"github.com/gorilla/mux"  // Import router Gorilla Mux
)

// Fungsi utama yang dijalankan saat aplikasi dimulai
func main() {
	// Inisialisasi router menggunakan Gorilla Mux
	// Router ini akan menangani semua routing HTTP
	r := mux.NewRouter()

	// Setup routing untuk aplikasi:
	// 1. Route utama ("/") untuk halaman web
	r.HandleFunc("/", handlers.WebHandler).Methods("GET")
	// 2. Route API ("/api/convert") untuk konversi suhu
	r.HandleFunc("/api/convert", handlers.APIHandler).Methods("GET")

	// Konfigurasi untuk menyajikan file statis
	// Semua file dalam folder 'static' akan tersedia di path '/static/'
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", 
			http.FileServer(http.Dir("static"))))

	// Menjalankan server HTTP
	log.Println("Server berjalan pada port 8080")
	// Server akan berjalan di port 8080 dan menggunakan router yang sudah dibuat
	// ListenAndServe bersifat blocking, akan return error jika terjadi masalah
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error: ", err)  // Log error dan exit jika server gagal berjalan
	}
}