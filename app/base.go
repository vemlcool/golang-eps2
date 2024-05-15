package app

import (
	"encoding/json"
	"encoding/xml"
)

type Base struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func (b *Base) SetIDAndName(id int, name string) {
	b.ID = id
	b.Name = name
}

func (b Base) MarshalJSON() ([]byte, error) {
	type Alias Base
	return json.Marshal((Alias)(b))
}

func (b Base) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Base
	return e.EncodeElement(Alias(b), start)
}
