package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

/*
XLM (eXtensible Markup Language) is used for encoding documents in a format that is both human readable and machine level
- Widely used for data enterchange between system and configurations files
- Go provides encoding/xml package for handling XML data
	- xml.Marshall (converts from GO Data Structures to XML format)
	- xml.Unmarshall (converts from XML format to GO Data Structures)
- Functions to encode and decode XML data
*/

type Person struct {
	// special field used by encoding/xml package and used to determine the name of xml element while marshalling and unmarshalling data
	// not to be set by user but to be automatically handled by encoding/xml package while handlaing marshalling and unmarshalling data
	XMLName xml.Name `xml:"person"` // root emelent will be <person> --> if no used root elment will be struct name i.e. <Person>
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city,omitempty"`
	State string `xml:"state,omitempty"`
}

func main() {

	person := Person{Name: "John", Age: 30, Address: Address{City: "New York", State: "NY"}, Email: "john@email.com"}

	// xmlData, err := xml.Marshal(person)
	xmlData, err := xml.MarshalIndent(person, "", "  ") // For better looking
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}
	fmt.Println(string(xmlData))

	rawXmlData := `<person><name>Jane</name><age>35</age><address><city>Kathmandu</city><state>Bagmati</state></address></person>`
	var personXml Person
	err = xml.Unmarshal([]byte(rawXmlData), &personXml)
	if err != nil {
		log.Fatalln("Error unmarshalling data:", err)
	}
	fmt.Println(personXml)
	fmt.Println("Local string:", personXml.XMLName.Local)
	fmt.Println("Namespace:", personXml.XMLName.Space)

	book := Book{
		ISBN:   "344-343-4343-34343",
		Title:  "Go BootCamp",
		Author: "Random",
		Pseudo: "Pesudo",
	}
	xmlAttrData, err := xml.MarshalIndent(book, "", "  ") // For better looking
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}
	fmt.Println(string(xmlAttrData))

}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"` // treating as attribute not child element i.e property
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Pseudo  string   `xml:"pseudo"`
}

/*
<book isbn="sadsadasda" color="blue">
</book>
*/
