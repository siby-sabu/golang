package lesson16

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice []string

func TestSlice() {

	var worldCupTeams = [...]string{"India", "Australia", "South Africa", "New Zealand", "Netherland"}
	var worldCupTeamsSliced = []string{"India", "Australia", "South Africa", "New Zealand", "Netherland"}

	//Any changes to a slice will also affect the underlying array
	finalsTeam := worldCupTeams[0:2]
	semiFinalsTeam := worldCupTeams[0:4]

	fmt.Println(worldCupTeams)
	fmt.Println(finalsTeam)
	fmt.Println(semiFinalsTeam)
	semiFinalsTeam[3] = "Pakistan"
	fmt.Println(worldCupTeams)
	fmt.Println(finalsTeam)
	fmt.Println(semiFinalsTeam)

	//Slicing a string
	helloWorld := "Hello World"
	world := helloWorld[6:]

	fmt.Println(helloWorld)
	fmt.Println(world)
	helloWorld = "Hello New World"

	fmt.Println(helloWorld)
	fmt.Println(world)

	fmt.Printf("%T\n", worldCupTeams)
	fmt.Printf("%T\n", worldCupTeamsSliced)

}

func removeSpaces(offices []string) {
	for i := range offices {
		offices[i] = strings.TrimSpace(offices[i])
	}
}

func TestSlicing() {

	officeLoc := []string{" Marathahalli ", "  KR Puram ", " Whitefield "}
	fmt.Printf("%v", strings.Join(officeLoc, "->"))
	removeSpaces(officeLoc)
	fmt.Printf("%v", strings.Join(officeLoc, "->"))
}

func passingArray(worldCup [4]string) {
	worldCup[3] = "Pakistan"
}

func passingSlice(worldCup []string) {
	worldCup[3] = "Pakistan"
}

func TestArrayAndSlice() {
	worldCupArray := [4]string{"India", "Australia", "New Zealand", "South Africa"}
	fmt.Println(worldCupArray)
	passingArray(worldCupArray) // No change since its sending a new copy of the array to the new function
	fmt.Println(worldCupArray)
	passingSlice(worldCupArray[:]) // Slice is just a view on top of the array. Any changes made through the slice will affect the array
	fmt.Println(worldCupArray)

}

func TestSliceSort() {

	worldCupArray := StringSlice{"India", "Austrlia", "NewZealand", "SouthAfrica"}
	fmt.Println(worldCupArray)
	worldCupArray.Sort()
	fmt.Println(worldCupArray)
}

func (str StringSlice) Sort() {
	sort.Slice(str, func(i, j int) bool {
		return str[i] > str[j]
	})
}

func AddingNewElementsToSlice() {

	/**In context of an array
	len is the length of the slice
	cap is the length of the underlying array

	**/
	worldCupTeams := []string{"India", "Australia", "New Zealand"}
	fmt.Println(worldCupTeams)
	//worldCupTeams = append(worldCupTeams, "South Africa")
	//fmt.Println(worldCupTeams)
	//finalTeams := worldCupTeams[0:2]
	finalTeams := worldCupTeams[0:2:2]
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeams), cap(worldCupTeams), worldCupTeams)
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(finalTeams), cap(finalTeams), finalTeams)
	finalTeams = append(finalTeams, "South Africa")
	//finalTeamWithCapacityRestricted := worldCupTeams[0:2:2] // Three index slicing where we mention the capacity pf the slice as thrid index
	//finalTeamWithCapacityRestricted = append(finalTeamWithCapacityRestricted, )
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeams), cap(worldCupTeams), worldCupTeams)
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(finalTeams), cap(finalTeams), finalTeams)

}

func PreAllocateWithMake() {
	worldCupTeams := []string{"India", "Australia"}
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeams), cap(worldCupTeams), worldCupTeams)
	worldCupTeams = append(worldCupTeams, "New Zealand")
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeams), cap(worldCupTeams), worldCupTeams)

	worldCupTeamsWithMake := make([]string, 8)
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeamsWithMake), cap(worldCupTeamsWithMake), worldCupTeamsWithMake)
	worldCupTeamsWithMake = append(worldCupTeamsWithMake, "South Africa")
	fmt.Printf("Length: %v, Capacity: %v, slice: %v \n", len(worldCupTeamsWithMake), cap(worldCupTeamsWithMake), worldCupTeamsWithMake)
	fmt.Println(worldCupTeamsWithMake[7])
	fmt.Println(worldCupTeamsWithMake[8])

}
