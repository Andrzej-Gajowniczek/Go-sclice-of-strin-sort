package main

import (
	"sort"
)

func ShowWhatsInside(pointer []string, comment string) {
	horizontalDivider(comment)
	for i, j := range pointer {
		println("position:", i, "\t item:", j)
	}

}
func horizontalDivider(paragraf string) {
	println("-------------------------------", paragraf, "-----------------------------------")
}

func main() {
	var fridge = []string{"butter", "milk", "becon", "sosages", "chees with holes", "red meat", "cream", "beef", "a dozen eggs"}
	ShowWhatsInside(fridge, "fridge")
	//do a copy of fridge
	copyOfFridge := make([]string, len(fridge))
	copy(copyOfFridge, fridge)

	ShowWhatsInside(copyOfFridge, "This is copy of fridge")

	// sort fridge but don't sort copy!!!
	sort.Strings(fridge)
	ShowWhatsInside(fridge, "this is sorted fridge")

	//Copy shouldn't be sorted because copy was made before sortin items

	ShowWhatsInside(copyOfFridge, "Copy of the fridge made before sorting")
}
