package lesson14

import (
	"fmt"
)

type Celsius float32
type Fahrenhiet float32

func (c Celsius) covertToFahrenhiet() Fahrenhiet {
	return Fahrenhiet((c * 9 / 5) + 32)
}

func (f Fahrenhiet) convertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func printRowSepereator() {
	fmt.Println("=========================================")
}

func printRow(col1 string, col2 string) {
	fmt.Printf("|\t%v\t|\t%v\t|\n", col1, col2)
}

func printHeader(head1 string, head2 string) {
	printRowSepereator()
	printRow(head1, head2)
	printRowSepereator()

}

func DrawCelsiusToFarenhietTable(start int, end int, step int) {
	printHeader("°C", "°F")
	for i := start; i <= end; i += step {
		c := Celsius(i)
		f := c.covertToFahrenhiet()
		printRow(fmt.Sprintf("%v", c), fmt.Sprintf("%v", f))

	}
}

func DrawFarenhietToCelsiusTable(start int, end int, step int) {
	printHeader("°F", "°C")
	for i := start; i <= end; i += step {
		f := Fahrenhiet(i)
		c := f.convertToCelsius()
		printRow(fmt.Sprintf("%v", f), fmt.Sprintf("%v", c))

	}
}
