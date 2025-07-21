package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// TemperatureResponse struktur untuk response API
// Berisi nilai suhu dalam tiga satuan yang berbeda
type TemperatureResponse struct {
	Celsius    float64 `json:"celsius"`    // Suhu dalam satuan Celsius
	Fahrenheit float64 `json:"fahrenheit"` // Suhu dalam satuan Fahrenheit
	Kelvin     float64 `json:"kelvin"`     // Suhu dalam satuan Kelvin
}

// APIHandler menangani request API konversi suhu
// Menerima parameter:
// - value: nilai suhu yang akan dikonversi (string)
// - from: satuan suhu asal ("celsius", "fahrenheit", atau "kelvin")
// Mengembalikan response JSON dengan suhu dalam ketiga satuan
func APIHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter dari URL
	valueStr := r.URL.Query().Get("value")
	fromUnit := r.URL.Query().Get("from")

	// Validasi input: konversi nilai suhu ke float64
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		respondWithError(w, "Invalid temperature value")
		return
	}

	// Variabel untuk menyimpan hasil konversi
	var celsius, fahrenheit, kelvin float64
	
	// Konversi suhu berdasarkan satuan asal
	switch fromUnit {
	case "celsius":
		// Jika asal Celsius, hitung Fahrenheit dan Kelvin
		celsius = value
		fahrenheit = (value * 9/5) + 32
		kelvin = value + 273.15
	case "fahrenheit":
		// Jika asal Fahrenheit, hitung Celsius dulu kemudian Kelvin
		celsius = (value - 32) * 5/9
		fahrenheit = value
		kelvin = celsius + 273.15
	case "kelvin":
		// Jika asal Kelvin, hitung Celsius dulu kemudian Fahrenheit
		celsius = value - 273.15
		fahrenheit = (celsius * 9/5) + 32
		kelvin = value
	default:
		// Jika satuan asal tidak valid
		respondWithError(w, "Invalid unit parameter")
		return
	}

	// Kirim response JSON dengan hasil konversi
	respondWithJSON(w, http.StatusOK, TemperatureResponse{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	})
}

// respondWithError helper function untuk mengirim response error
// Parameters:
// - w: http.ResponseWriter untuk menulis response
// - message: pesan error yang akan dikirim
func respondWithError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// respondWithJSON helper function untuk mengirim response JSON
// Parameters:
// - w: http.ResponseWriter untuk menulis response
// - code: HTTP status code
// - payload: data yang akan diencode ke JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}