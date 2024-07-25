package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/go-dependency-injection/printer"
	"github.com/OlyMahmudMugdho/go-dependency-injection/writer"
)

func main() {

	redWriter := writer.NewWriter(&printer.RedPrinter{})
	fmt.Println(redWriter.Write())

	blackWriter := writer.NewWriter(&printer.BlackPrinter{})
	fmt.Println(blackWriter.Write())

	blueWriter := writer.NewWriter(&printer.BluePrinter{})
	fmt.Println(blueWriter.Write())

}
