package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	data, err := ioutil.ReadFile("notes.xml") //File read

	if err != nil {
		fmt.Println("Unable to open file")
		return
	}

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), &note) //Unmarshalling bytes [] into structure...

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)

	xmlwriter()
}

func xmlwriter() {

	note := Notes{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}

	file, _ := xml.MarshalIndent(note, "", " ")

	_ = ioutil.WriteFile("notes1.xml", file, 0644)
}
