package mplus

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-netsuite-soap/omitempty"
)

type Customer struct {
}

func (c Customer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(c, e, start)
}

func (c Customer) IsEmpty() bool {
	return zero.IsZero(c)
}
