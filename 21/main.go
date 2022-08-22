package main

import "fmt"

/*
TODO Реализовать паттерн «адаптер» на любом примере.
*/

type AnaliticService interface {
	SendXml() interface{}
}

type XmlDocument struct {
}

func (doc *XmlDocument) SendXml() {
	fmt.Println("sending XML to service")
}

type JsonDocument struct {
}

func (doc *JsonDocument) ConvertToXml() string {
	return ""
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter *JsonDocumentAdapter) SendXml() {
	adapter.jsonDocument.ConvertToXml()
	fmt.Println("sending data to analitic service")
}

func main() {
	xmlFile := &XmlDocument{}
	xmlFile.SendXml()

	jsonFile := &JsonDocument{}
	adapter := JsonDocumentAdapter{jsonDocument: jsonFile}
	adapter.SendXml()
}
