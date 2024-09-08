package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Initialize the collector
	c := colly.NewCollector()

	// slice to hold weather data
	var weatherData [][]string

	// Scrape weather information for morning
	c.OnHTML("tr.morning", func(e *colly.HTMLElement) {
		day := e.ChildText("td.weather-day")
		temp := e.ChildText("td.weather-temperature span")
		if day != "" && temp != "" {
			fmt.Printf("Day: %s, Temperature: %s\n", day, temp)
			weatherData = append(weatherData, []string{day, temp})
		}
	})

	// scraping for night
	c.OnHTML("tr.night", func(e *colly.HTMLElement) {
		day := e.ChildText("td.weather-day")
		temp := e.ChildText("td.weather-temperature span")
		if day != "" && temp != "" {
			fmt.Printf("Day: %s, Temperature: %s\n", day, temp)
			weatherData = append(weatherData, []string{day, temp})
		}
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://world-weather.info/forecast/usa/orlando/")

	// collected data to CSV
	writeCSV(weatherData)
}
