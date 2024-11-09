package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ConversionData struct {
	Value      float64
	FromUnit   string
	ToUnit     string
	Result     float64
	Error      string
	ShowResult bool
}

type ConversionMap map[string]float64

var (
	lengthConversions = ConversionMap{
		"millimeter": 0.001,
		"centimeter": 0.01,
		"meter":      1,
		"kilometer":  1000,
		"inch":       0.0254,
		"foot":       0.3048,
		"yard":       0.9144,
		"mile":       1609.344,
	}

	weightConversions = ConversionMap{
		"milligram": 0.001,
		"gram":      1,
		"kilogram":  1000,
		"ounce":     28.349523125,
		"pound":     453.59237,
	}

	temperatureUnits = []string{"celsius", "fahrenheit", "kelvin"}
)

type Converter interface {
	Convert(value float64, from, to string) float64
}

type LengthConverter struct {
	conversions ConversionMap
}

type WeightConverter struct {
	conversions ConversionMap
}

type TemperatureConverter struct{}

func (lc LengthConverter) Convert(value float64, from, to string) float64 {
	baseValue := value * lc.conversions[from]
	return baseValue / lc.conversions[to]
}

func (wc WeightConverter) Convert(value float64, from, to string) float64 {
	baseValue := value * wc.conversions[from]
	return baseValue / wc.conversions[to]
}

func (tc TemperatureConverter) Convert(value float64, from, to string) float64 {
	celsiusTemp := value
	switch from {
	case "fahrenheit":
		celsiusTemp = (value - 32) * 5 / 9
	case "kelvin":
		celsiusTemp = value - 273.15
	}

	switch to {
	case "celsius":
		return celsiusTemp
	case "fahrenheit":
		return (celsiusTemp * 9 / 5) + 32
	case "kelvin":
		return celsiusTemp + 273.15
	}
	return 0
}

type Server struct {
	converters map[string]Converter
	templates  *template.Template
}

func NewServer() *Server {
	return &Server{
		converters: map[string]Converter{
			"length":      LengthConverter{conversions: lengthConversions},
			"weight":      WeightConverter{conversions: weightConversions},
			"temperature": TemperatureConverter{},
		},
		templates: template.Must(template.ParseFiles("templates/index.html", "templates/converter.html")),
	}
}

func (s *Server) handleConversion(w http.ResponseWriter, r *http.Request, convType string) {
	data := ConversionData{ShowResult: false}

	if r.Method == http.MethodPost {
		value, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			data.Error = "Invalid number entered"
			data.ShowResult = true
		} else {
			data.Value = value
			data.FromUnit = r.FormValue("fromUnit")
			data.ToUnit = r.FormValue("toUnit")

			converter, ok := s.converters[convType]
			if !ok {
				http.Error(w, "Invalid conversion type", http.StatusBadRequest)
				return
			}

			data.Result = converter.Convert(value, data.FromUnit, data.ToUnit)
			data.ShowResult = true
		}
	}

	err := s.templates.ExecuteTemplate(w, "converter.html", struct {
		Type  string
		Data  ConversionData
		Units interface{}
	}{
		Type:  convType,
		Data:  data,
		Units: getUnits(convType),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUnits(convType string) []string {
	switch convType {
	case "length":
		return []string{"millimeter", "centimeter", "meter", "kilometer", "inch", "foot", "yard", "mile"}
	case "weight":
		return []string{"milligram", "gram", "kilogram", "ounce", "pound"}
	case "temperature":
		return temperatureUnits
	}
	return []string{}
}

func main() {
	server := NewServer()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.templates.ExecuteTemplate(w, "index.html", nil)
	})

	// Create a single handler for all conversion types
	conversionHandler := func(convType string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			server.handleConversion(w, r, convType)
		}
	}

	// Register routes using the conversionHandler
	http.HandleFunc("/length", conversionHandler("length"))
	http.HandleFunc("/weight", conversionHandler("weight"))
	http.HandleFunc("/temperature", conversionHandler("temperature"))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
