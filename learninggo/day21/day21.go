package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	// Đổi sang MethodGet để gọi được từ trình duyệt/Postman
	if r.Method != http.MethodGet {
		http.Error(w, "Phuong thuc nay khong duoc ho tro", http.StatusMethodNotAllowed)
		return
	}

	// Đổi tên biến đồng nhất thành response
	response := map[string]string{
		"message": "Chao mung cac ban den voi khoa hoc lap trinh Golang",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Course", "Lap trinh Golang")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Loi ma hoa JSON", http.StatusInternalServerError)
		return
	}
}
