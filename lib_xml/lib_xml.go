package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
    Name  string `xml:"name"`
    Age   int    `xml:"age"`
    City  string `xml:"city"`
}

func main() {
    person := Person{
        Name:  "John",
        Age:   30,
        City:  "New York",
    }

    xmlData, _ := xml.Marshal(person)
    fmt.Println("XML Data:", string(xmlData))
}

/*
$ go run .
XML Data: <Person><name>John</name><age>30</age><city>New York</city></Person>
*/