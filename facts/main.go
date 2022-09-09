package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Tris console app can bring you fact about any number!")

	var number int
	var ftype string

	var factTypes = map[string]string{
		"trivia": "Факт из жизни",
		"math":   "Факт из математики",
	}

	fmt.Print("Input a number: ")
	fmt.Fscan(os.Stdin, &number)

	fmt.Println("Fact types: ")
	for key, value := range factTypes {
		fmt.Println(key + " - " + value)
	}

	fmt.Print("Input a type: ")
	fmt.Fscan(os.Stdin, &ftype)

	url := "http://numbersapi.com/" + strconv.Itoa(number) + "/" + ftype

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println("Fact: " + sb)
}
