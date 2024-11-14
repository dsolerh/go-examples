package main

import "fmt"

func main() {
	// teams := [][]string{
	// 	{"A", "B", "C"},
	// 	{"D", "E"},
	// }

	// // [A D B E C nil]

	// maxTeamLengt := 3
	// playerOrder := make([]*string, maxTeamLengt*len(teams))
	// for i, team := range teams {
	// 	for j := range team {
	// 		member := team[j]
	// 		playerOrder[j*len(teams)+i] = &member
	// 	}
	// }

	// fmt.Printf("playerOrder: %v\n", playerOrder)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("a[:5]: %v\n", a[:5])
	fmt.Printf("a[5:]: %v\n", a[5:])

	a = nil
	a = append(a, []int{1, 4, 7}...)
	fmt.Printf("a: %v\n", a)
}
