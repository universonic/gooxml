// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_FtnEdnSepRef struct {
	// Footnote/Endnote ID
	IdAttr int64
}

func NewCT_FtnEdnSepRef() *CT_FtnEdnSepRef {
	ret := &CT_FtnEdnSepRef{}
	return ret
}

func (m *CT_FtnEdnSepRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FtnEdnSepRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FtnEdnSepRef: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FtnEdnSepRef and its children
func (m *CT_FtnEdnSepRef) Validate() error {
	return m.ValidateWithPath("CT_FtnEdnSepRef")
}

// ValidateWithPath validates the CT_FtnEdnSepRef and its children, prefixing error messages with path
func (m *CT_FtnEdnSepRef) ValidateWithPath(path string) error {
	return nil
}
