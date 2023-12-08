package lesson16

import "fmt"

func updateOffice(offices [8]string) {
	offices[2] = "MG Road"
}

func Test() {
	// var planets [8]string
	// planets[0] = "Mercury"
	// planets[1] = "Venus"
	// fmt.Println(planets[1])
	// fmt.Println(len(planets))
	//planets[10] = "Dumy" //compile time error

	// i := 10
	// planets[i] = "Pluto" //Panic error

	//composite literals
	// var bangaloreOffices = [3]string{"whitefield", "marathahalli", "CV raman nagar"}
	// fmt.Print(bangaloreOffices)

	bangaloreOffices := [...]string{"whitefield", "marathahalli", "CV raman nagar"}
	//fmt.Print(len(bangaloreOffices))

	//_ represents blank variable
	//Iterating an array
	// for pos, office := range bangaloreOffices {
	// 	fmt.Printf("%v : %v \n", pos, office)

	// }

	bangaloreNewOffices := bangaloreOffices
	bangaloreNewOffices[2] = "NR Enclave"
	// fmt.Println(bangaloreNewOffices)
	// fmt.Println(bangaloreOffices)

	//updateOffice(bangaloreOffices)
	fmt.Println(bangaloreOffices)

}
