package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Conversion struct {
	Value  float64
	From   string
	To     string
	Result float64
	Type   string
}

// Conversion factors for length (to meters)
var lengthFactors = map[string]float64{
	// standard
	"millimeter": 0.001,
	"centimeter": 0.01,
	"meter":      1.0,
	"kilometer":  1000.0,

	// empirical
	"inch": 0.0254,
	"foot": 0.3048,
	"yard": 0.9144,
	"mile": 1609.34,
}

// Conversion factors for weight (to kilograms)
var weightFactors = map[string]float64{
	// standard
	"milligram": 0.000001,
	"gram":      0.001,
	"kilogram":  1.0,

	// empirical
	"ounce": 0.0283495,
	"pound": 0.453592,
}

func convertLength(value float64, from, to string) float64 {
	meters := value * lengthFactors[from]
	return meters / lengthFactors[to]
}

func convertWeight(value float64, from, to string) float64 {
	kilograms := value * weightFactors[from]
	return kilograms / weightFactors[to]
}

func convertTemperature(value float64, from, to string) float64 {
	// Convert to Kelvin first
	var kelvin float64
	switch from {
	case "celsius":
		kelvin = value + 273.15
	case "fahrenheit":
		kelvin = (value + 459.67) * 5 / 9
	case "kelvin":
		kelvin = value
	}

	// Convert from Kelvin to target unit
	switch to {
	case "celsius":
		return kelvin - 273.15
	case "fahrenheit":
		return kelvin*9/5 - 459.67
	case "kelvin":
		return kelvin
	}
	return 0
}

func handleConversion(w http.ResponseWriter, r *http.Request) {
	convType := r.URL.Path[1:] // Get conversion type from URL
	if convType == "" {
		convType = "length"
	}

	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	conv := Conversion{Type: convType}

	if r.Method == http.MethodPost {
		value, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err == nil {
			conv.Value = value
			conv.From = r.FormValue("from")
			conv.To = r.FormValue("to")

			switch convType {
			case "length":
				conv.Result = convertLength(value, conv.From, conv.To)
			case "weight":
				conv.Result = convertWeight(value, conv.From, conv.To)
			case "temperature":
				conv.Result = convertTemperature(value, conv.From, conv.To)
			}
		}
	}

	_ = tmpl.Execute(w, conv)
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", handleConversion)
	http.HandleFunc("/length", handleConversion)
	http.HandleFunc("/weight", handleConversion)
	http.HandleFunc("/temperature", handleConversion)

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
