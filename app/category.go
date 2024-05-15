package app

import "encoding/xml"

type Category struct {
	Base
	XMLName xml.Name `json:"-"`
}
