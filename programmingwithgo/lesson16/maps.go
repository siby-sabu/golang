package lesson16

import "fmt"

func BasicMaps() {
	topScores := map[string]int{
		"South Africa": 428,
		"India":        410,
		"New Zealand":  401,
	}

	fmt.Printf("Highest score by India in 2023 WC was %v", topScores["India"])
	country := "Australia"
	if score, found := topScores[country]; found {
		fmt.Printf("Top score from %v is %v", country, score)
	} else {
		fmt.Printf("%v is not there is top scored team list", country)
	}
}

func AQI() {
	aqiByCities := map[string]int{
		"Delhi":       423,
		"Ulaanbaatar": 183,
		"Kolkata":     182,
		"Karachi":     181,
	}

	fmt.Println(aqiByCities)
	for country, aqi := range aqiByCities {
		fmt.Printf("Air Quality Index of %v is %v \n", country, aqi)

	}
	if val, ok := aqiByCities["Delhi"]; ok {
		fmt.Printf("%v is found in AQI list", val)
	} else {
		fmt.Printf("%v is not found in AQI list", val)
	}
}

func AssignValues() {

	topScores := map[string]int{
		"South Africa": 428,
		"India":        410,
		"New Zealand":  401,
	}
	expectedTopScores := topScores
	expectedTopScores["India"] = 450

	fmt.Printf("%v \n", topScores)
	delete(expectedTopScores, "South Africa")
	fmt.Printf("%v \n", topScores)

}
