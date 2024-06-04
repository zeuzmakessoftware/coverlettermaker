package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func checkStringWithWordList(string_thing string, word_list []string, company string, position string, name string) []string {
	if string_thing[1:] == "[company]" {
		word_list = append(word_list, company)
	} else if string_thing[1:] == "[position]" {
		word_list = append(word_list, position)
	} else if string_thing[1:] == "[name]" {
		word_list = append(word_list, name)
	} else {
		word_list = append(word_list, string_thing)
	}
	return word_list
}

func generatePDF(content string, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.MultiCell(190, 10, content, "", "", false)
	return pdf.OutputFileAndClose(filename)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		fmt.Println("go run CLMaker.go [company] [position] [name] [optional: -pdf]")
		fmt.Println("./CLMaker [company] [position] [name] [optional: -pdf]")
		return
	}
	if len(os.Args) < 4 {
		fmt.Println("Please provide more arguments")
		return
	}
	if len(os.Args) > 4 && os.Args[4] != "-pdf" {
		fmt.Println("Too many arguments, are you surrounding with quotes?")
		return
	}

	company := fmt.Sprintf(" %s", os.Args[1])
	position := fmt.Sprintf(" %s", os.Args[2])
	name := fmt.Sprintf(" %s", os.Args[3])
	outputToPDF := len(os.Args) == 5 && os.Args[4] == "-pdf"

	file, err := os.ReadFile("coverletterexample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	word_list := []string{}
	string_thing := ""
	for _, data := range file {
		if string(data) == "," || string(data) == "." || string(data) == "\n" {
			word_list = checkStringWithWordList(string_thing, word_list, company, position, name)
			string_thing = ""
		}
		if string(data) == " " {
			word_list = checkStringWithWordList(string_thing, word_list, company, position, name)
			string_thing = " "
		} else {
			string_thing += string(data)
		}
	}
	word_list = checkStringWithWordList(string_thing, word_list, company, position, name)
	string_thing = " "
	content := strings.Join(word_list, "")

	if outputToPDF {
		err = generatePDF(content, "cover_letter.pdf")
		if err != nil {
			fmt.Println("Error generating PDF:", err)
			return
		}
		fmt.Println("Output written to cover_letter.pdf")
	} else {
		fmt.Println(content)
	}
}
