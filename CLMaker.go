package main

import (
	"fmt"
	"os"
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

func main() {
	if os.Args[1] == "help" {
		fmt.Println("go run CLMaker.go [company] [position] [name]")
		fmt.Println("or")
		fmt.Println("./CLMaker [company] [position] [name]")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}
	if len(os.Args) < 4 {
		fmt.Println("Please provide more arguments")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("Too many arguments, are you surrounding with quotes?")
		return
	}
	company := fmt.Sprintf(" %s", os.Args[1])
	position := fmt.Sprintf(" %s", os.Args[2])
	name := fmt.Sprintf(" %s", os.Args[3])
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
	for _, word := range word_list {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}
