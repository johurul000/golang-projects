package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jung-kurt/gofpdf"
)
var file string
var pdfFile string

func main(){

	getFile()

	content, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	createPdf(string(content))

	
}

func getFile(){
	fmt.Println("Enter the name of the txt file( '.txt' is a must): ")
	fmt.Scan(&file)
	fmt.Println("Enter the name of the pdf file to be Generated( '.pdf' is amust): ")
	fmt.Scan(&pdfFile)
}

func createPdf(content string){
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose(pdfFile)

	fmt.Println("PDF Created")

}