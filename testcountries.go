package main

import "fmt"
import "github.com/biter777/countries"

func main() {
	countryJapan := countries.Japan
	fmt.Printf("Country name: %v\n", countryJapan)
	fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())
	fmt.Printf("Country ISO2: %v\n", countryJapan.ISO2())
	fmt.Printf("Country ISO3: %v\n", countryJapan.ISO3())
}
