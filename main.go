package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

func ShowCoparison(pointerA []string, pointerB []string, comment string) {
	horizontalDivider(comment)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 24, 40, 0, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, " %s\t%s\t%s\t\n", "index", "fridge", "fridge-copy")
	fmt.Fprintf(w, " %s\t%s\t%s\t\n", "----", "----", "----")

	for i, j := range pointerA {

		fmt.Fprintf(w, " %v\t%v\t%v\t\n", i, j, pointerB[i])

	}

}

func ShowWhatsInside(pointer []string, comment string) {
	horizontalDivider(comment)
	for i, j := range pointer {
		println("position:", i, "\t item:", j)
	}

}
func horizontalDivider(paragraf string) {
	maxWidth := 80
	texWidth := len(paragraf)
	tailWidth := (maxWidth - texWidth) / 2
	for i := 0; i < maxWidth; i++ {
		if i < tailWidth || i > (tailWidth+texWidth) {
			fmt.Print("-")

		} else {
			i = i + texWidth
			fmt.Print(paragraf)
		}
	}

	println("")
}

func main() {
	var fridge = []string{"butter", "milk", "becon", "sausages", "chees with holes", "red meat", "cream", "beef", "dozen eggs"}
	ShowWhatsInside(fridge, "fridge")
	//do a copy of fridge
	copyOfFridge := make([]string, len(fridge))
	copy(copyOfFridge, fridge)

	ShowWhatsInside(copyOfFridge, "This is copy of fridge")

	// sort fridge but don't sort copy!!!
	sort.Strings(fridge)

	ShowCoparison(fridge, copyOfFridge, "The cmparison of the fridge and the copy")
	horizontalDivider("")
}
