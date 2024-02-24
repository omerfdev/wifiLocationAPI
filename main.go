package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	City string `json:"city"`
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	// IP adresini al
	ip := r.RemoteAddr

	// IP adresine göre konum bilgisini almak için bir istek oluştur
	url := fmt.Sprintf("https://ipinfo.io/%s?token=YOUR_API_TOKEN", ip)

	// API'ye GET isteği yap
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// API yanıtını oku
	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON formatında konum bilgilerini döndür
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func main() {
	http.HandleFunc("/location", getLocation)

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
