package app

import "encoding/xml"

type Product struct {
	Base
	Price    int64 `json:"price" xml:"price"`
	Year     int
	Category Category `json:"category" xml:"category"`
	XMLName  xml.Name `json:"-"`
}

func (p *Product) GetYear() int {
	return p.Year
}

func (p *Product) SetYear(Year int) {
	p.Year = Year
}
