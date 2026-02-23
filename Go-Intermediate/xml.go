package main

import (
	"encoding/xml"
	"fmt"
)

type persontype struct {
	XMLName xml.Name `xml:"personXML"`
	Name    string   `xml:"name,omitempty"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,attr"`
}

func xmlExample() {
	person := persontype{
		Name:  "Nirmit Sahu",
		Age:   30,
		Email: "john@example.com",
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println(string(xmlData))

	xmlData2, _ := xml.MarshalIndent(person, "", "  ")
	fmt.Println(string(xmlData2))

	// Unmarshalling
	xmlInput := `<personXML email="john@example.com">
		<name>Nirmit Sahu</name>
		<age>30</age>
	</personXML>`
	var p persontype
	err = xml.Unmarshal([]byte(xmlInput), &p)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
	fmt.Printf("Unmarshalled Person: %+v\n", p)
}
