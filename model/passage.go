package model

import "encoding/xml"

type Verse struct {
	XMLName xml.Name `xml:"verse"`
	Number  int32    `xml:"number"`
	Title   string   `xml:"title,omitempty"`
	Text    string   `xml:"text"`
}

type Verses struct {
	XMLName xml.Name `xml:"verses"`
	Verse   []Verse  `xml:"verse"`
}

type Chapter struct {
	XMLName xml.Name `xml:"chapter"`
	Chap    int32    `xml:"chap"`
	Verses  Verses   `xml:"verses"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
	Chapter Chapter  `xml:"chapter"`
}

type Passage struct {
	XMLName xml.Name `xml:"bible"`
	Book    Book     `xml:"book"`
}
