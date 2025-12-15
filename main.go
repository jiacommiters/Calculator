package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	Result string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Tampilkan form kalkulator
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Data kosong untuk pertama kali
	data := PageData{
		Result: "",
	}

	tmpl.Execute(w, data)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse form
	r.ParseForm()
	
	// Ambil data dari form
	expression := r.FormValue("expression")
	result := r.FormValue("result")

	fmt.Printf("Received - Expression: %s, Result: %s\n", expression, result)

	// Jika ada result langsung, gunakan itu
	if result != "" {
		renderResult(w, result)
		return
	}

	// Jika tidak, hitung dari expression
	if expression != "" {
		calculatedResult := calculateExpression(expression)
		renderResult(w, calculatedResult)
		return
	}

	// Jika tidak ada data, redirect ke home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func calculateExpression(expr string) string {
	// Bersihkan expression
	expr = strings.ReplaceAll(expr, "×", "*")
	expr = strings.ReplaceAll(expr, "÷", "/")
	expr = strings.ReplaceAll(expr, " ", "")

	// Parsing sederhana untuk operasi dasar
	// Catatan: Ini adalah evaluasi sederhana, 
	// untuk produksi gunakan parser matematika yang aman
	
	// Cari operator
	var result float64
	if strings.Contains(expr, "+") {
		parts := strings.Split(expr, "+")
		if len(parts) == 2 {
			a, _ := strconv.ParseFloat(parts[0], 64)
			b, _ := strconv.ParseFloat(parts[1], 64)
			result = a + b
		}
	} else if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) == 2 {
			a, _ := strconv.ParseFloat(parts[0], 64)
			b, _ := strconv.ParseFloat(parts[1], 64)
			result = a - b
		}
	} else if strings.Contains(expr, "*") {
		parts := strings.Split(expr, "*")
		if len(parts) == 2 {
			a, _ := strconv.ParseFloat(parts[0], 64)
			b, _ := strconv.ParseFloat(parts[1], 64)
			result = a * b
		}
	} else if strings.Contains(expr, "/") {
		parts := strings.Split(expr, "/")
		if len(parts) == 2 {
			a, _ := strconv.ParseFloat(parts[0], 64)
			b, _ := strconv.ParseFloat(parts[1], 64)
			if b != 0 {
				result = a / b
			} else {
				return "Error: Division by zero"
			}
		}
	} else {
		// Jika tidak ada operator, kembalikan angka asli
		return expr
	}

	return fmt.Sprintf("%.2f", result)
}

func renderResult(w http.ResponseWriter, result string) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: result,
	}

	tmpl.Execute(w, data)
}

func main() {
	// Handle routes
	http.HandleFunc("/", handler)
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Server Listening on :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	http.ListenAndServe(":8080", nil)
}